/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controllers

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	cloudnativepg "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/cloudnative-pg"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/aws"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/shardingsphere"

	cnpg "github.com/cloudnative-pg/cloudnative-pg/api/v1"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/strings/slices"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	StorageNodeControllerName = "storage-node-controller"
	FinalizerName             = "shardingsphere.apache.org/finalizer"

	AnnotationKeyRegisterStorageUnitEnabled = "shardingsphere.apache.org/register-storage-unit-enabled"
	AnnotationKeyComputeNodeName            = "shardingsphere.apache.org/compute-node-name"
	AnnotationKeyLogicDatabaseName          = "shardingsphere.apache.org/logic-database-name"

	ShardingSphereProtocolType = "proxy-frontend-database-protocol-type"
)

// StorageNodeReconciler is a controller for storage nodes
type StorageNodeReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Log      logr.Logger
	Recorder record.EventRecorder
	AwsRDS   rds.RDS
	CNPG     cloudnativepg.CloudNativePG

	Service service.Service
}

// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=storagenodes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=storagenodes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=storagenodes/finalizers,verbs=update
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=storageproviders,verbs=get;list;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=clusters,verbs=get;list;watch;create;update;patch;delete

// Reconcile handles main function of this controller
// nolint:gocognit
func (r *StorageNodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info(fmt.Sprintf("Reconciling StorageNode %s", req.NamespacedName))

	// get storage node
	node := &v1alpha1.StorageNode{}
	if err := r.Get(ctx, req.NamespacedName, node); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Get storageProvider with storagenode.Spec.StorageProviderName
	storageProvider, err := r.getStorageProvider(ctx, node)
	if err != nil {
		r.Log.Error(err, fmt.Sprintf("unable to fetch storageProvider %s", node.Spec.StorageProviderName))
		return ctrl.Result{Requeue: true}, err
	}

	// finalize storage node
	if node.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent to registering our finalizer.
		if !slices.Contains(node.ObjectMeta.Finalizers, FinalizerName) {
			node.ObjectMeta.Finalizers = append(node.ObjectMeta.Finalizers, FinalizerName)
			if err := r.Update(ctx, node); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else if slices.Contains(node.ObjectMeta.Finalizers, FinalizerName) {
		return r.finalize(ctx, node, storageProvider)
	}

	// reconcile storage node
	return r.reconcile(ctx, storageProvider, node)
}

func (r *StorageNodeReconciler) finalize(ctx context.Context, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) (ctrl.Result, error) {
	var err error
	var oldStatus = node.Status.DeepCopy()

	switch node.Status.Phase {
	case v1alpha1.StorageNodePhaseReady, v1alpha1.StorageNodePhaseNotReady:
		// set storage node status to deleting
		node.Status.Phase = v1alpha1.StorageNodePhaseDeleting
	case v1alpha1.StorageNodePhaseDeleting:
		break
	case v1alpha1.StorageNodePhaseDeleteComplete:
		node.ObjectMeta.Finalizers = slices.Filter([]string{}, node.ObjectMeta.Finalizers, func(f string) bool {
			return f != FinalizerName
		})
		if err = r.Update(ctx, node); err != nil {
			r.Log.Error(err, "failed to remove finalizer")
		}
		return ctrl.Result{}, nil
	}

	// Try to unregister storage unit in shardingsphere.
	if err = r.unregisterStorageUnit(ctx, node); err != nil {
		r.Log.Error(err, "failed to delete storage unit")
		return ctrl.Result{RequeueAfter: defaultRequeueTime}, err
	}

	if err = r.deleteDatabaseCluster(ctx, node, storageProvider); err != nil {
		r.Log.Error(err, "failed to delete database cluster")
		return ctrl.Result{RequeueAfter: defaultRequeueTime}, err
	}

	desiredState := computeDesiredState(node.Status)

	if !reflect.DeepEqual(oldStatus, desiredState) {
		node.Status = desiredState
		err := r.Status().Update(ctx, node)
		if err != nil {
			r.Log.Error(err, fmt.Sprintf("unable to update StorageNode %s/%s", node.GetNamespace(), node.GetName()))
			return ctrl.Result{Requeue: true}, err
		}
	}

	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
}

func (r *StorageNodeReconciler) reconcile(ctx context.Context, storageProvider *v1alpha1.StorageProvider, node *v1alpha1.StorageNode) (ctrl.Result, error) {
	var err error
	var oldStatus = node.Status.DeepCopy()

	// reconcile storage node with storageProvider
	switch storageProvider.Spec.Provisioner {
	case v1alpha1.ProvisionerAWSRDSInstance:
		if err := r.reconcileAwsRdsInstance(ctx, aws.NewRdsClient(r.AwsRDS), node, storageProvider); err != nil {
			r.Recorder.Eventf(node, corev1.EventTypeWarning, "Reconcile Failed", fmt.Sprintf("unable to reconcile AWS RDS Instance %s/%s, err:%s", node.GetNamespace(), node.GetName(), err.Error()))
			return ctrl.Result{RequeueAfter: defaultRequeueTime}, err
		}
	case v1alpha1.ProvisionerAWSRDSCluster:
		if err := r.reconcileAwsRDSCluster(ctx, aws.NewRdsClient(r.AwsRDS), node, storageProvider); err != nil {
			r.Recorder.Eventf(node, corev1.EventTypeWarning, "Reconcile Failed", fmt.Sprintf("unable to reconcile AWS RDS Cluster %s/%s, err:%s", node.GetNamespace(), node.GetName(), err.Error()))
			return ctrl.Result{RequeueAfter: defaultRequeueTime}, err
		}
	case v1alpha1.ProvisionerAWSAurora:
		if err := r.reconcileAwsAurora(ctx, aws.NewRdsClient(r.AwsRDS), node, storageProvider); err != nil {
			r.Recorder.Eventf(node, corev1.EventTypeWarning, "Reconcile Failed", fmt.Sprintf("unable to reconcile AWS Aurora %s/%s, err:%s", node.GetNamespace(), node.GetName(), err.Error()))
			return ctrl.Result{RequeueAfter: defaultRequeueTime}, err
		}
	case v1alpha1.ProvisionerCloudNativePG:
		if err := r.reconcileCloudNativePG(ctx, node, storageProvider); err != nil {
			r.Recorder.Eventf(node, corev1.EventTypeWarning, "Reconcile Failed", fmt.Sprintf("unable to reconcile CloudNative PG %s/%s, err:%s", node.GetNamespace(), node.GetName(), err.Error()))
			return ctrl.Result{RequeueAfter: defaultRequeueTime}, err
		}
	default:
		r.Recorder.Event(node, corev1.EventTypeWarning, "UnsupportedDatabaseProvisioner", fmt.Sprintf("unsupported database provisioner %s", storageProvider.Spec.Provisioner))
		return ctrl.Result{RequeueAfter: defaultRequeueTime}, err
	}

	// register storage unit if needed.
	if err := r.registerStorageUnit(ctx, node, storageProvider); err != nil {
		r.Recorder.Eventf(node, corev1.EventTypeWarning, "RegisterStorageUnitFailed", "unable to register storage unit %s/%s", node.GetNamespace(), node.GetName())
		return ctrl.Result{Requeue: true}, err
	}

	desiredState := computeDesiredState(node.Status)

	if !reflect.DeepEqual(oldStatus, desiredState) {
		node.Status = desiredState
		err := r.Status().Update(ctx, node)
		if err != nil {
			r.Log.Error(err, fmt.Sprintf("unable to update StorageNode %s/%s", node.GetNamespace(), node.GetName()))
			return ctrl.Result{Requeue: true}, err
		}
	}

	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
}

func (r *StorageNodeReconciler) getStorageProvider(ctx context.Context, node *v1alpha1.StorageNode) (storageProvider *v1alpha1.StorageProvider, err error) {
	if node.Spec.StorageProviderName == "" {
		r.Recorder.Event(node, corev1.EventTypeWarning, "storageProviderNameIsNil", "storageProviderName is nil")
		return nil, fmt.Errorf("storageProviderName is nil")
	}

	storageProvider = &v1alpha1.StorageProvider{}

	if err := r.Get(ctx, client.ObjectKey{Name: node.Spec.StorageProviderName}, storageProvider); err != nil {
		r.Log.Error(err, fmt.Sprintf("unable to fetch storageProvider %s", node.Spec.StorageProviderName))
		r.Recorder.Event(node, corev1.EventTypeWarning, "storageProviderNotFound", fmt.Sprintf("storageProvider %s not found", node.Spec.StorageProviderName))
		return nil, err
	}

	// check provisioner
	// aws-like provisioner need aws rds client
	if storageProvider.Spec.Provisioner == v1alpha1.ProvisionerAWSRDSInstance ||
		storageProvider.Spec.Provisioner == v1alpha1.ProvisionerAWSAurora ||
		storageProvider.Spec.Provisioner == v1alpha1.ProvisionerAWSRDSCluster {
		if r.AwsRDS == nil {
			r.Recorder.Event(node, corev1.EventTypeWarning, "AwsRdsClientIsNil", "aws rds client is nil, please check your aws credentials")
			return nil, fmt.Errorf("aws rds client is nil, please check your aws credentials")
		}
	}

	return storageProvider, nil
}

// nolint:gocritic
func computeDesiredState(status v1alpha1.StorageNodeStatus) v1alpha1.StorageNodeStatus {
	// Initialize a new status object based on the current state
	desiredState := status
	clusterStatus := status.Cluster.Status

	if status.Phase == v1alpha1.StorageNodePhaseDeleting {
		// If the storage node is being deleted, check if all instances are deleted.
		if clusterStatus == "" && len(status.Instances) == 0 {
			desiredState.Phase = v1alpha1.StorageNodePhaseDeleteComplete
		}
	} else {
		// If the storage node is not being deleted, check if all instances are ready.
		if (clusterStatus == "" || clusterStatus == string(rds.DBClusterStatusAvailable)) && allInstancesReady(status.Instances) {
			desiredState.Phase = v1alpha1.StorageNodePhaseReady
		} else {
			desiredState.Phase = v1alpha1.StorageNodePhaseNotReady
		}
	}

	desiredState.Conditions = computeNewConditions(desiredState, status, clusterStatus)

	return desiredState
}

// nolint:gocritic
func computeNewConditions(desiredState, status v1alpha1.StorageNodeStatus, clusterStatus string) v1alpha1.StorageNodeConditions {
	newSNConditions := status.Conditions

	// Update the cluster ready condition if the cluster status is not empty
	if clusterStatus != "" {
		if clusterStatus == string(rds.DBClusterStatusAvailable) {
			newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
				Type:           v1alpha1.StorageNodeConditionTypeClusterReady,
				Status:         corev1.ConditionTrue,
				LastUpdateTime: metav1.Now(),
				Reason:         "Cluster is ready",
			})
		} else {
			newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
				Type:           v1alpha1.StorageNodeConditionTypeClusterReady,
				Status:         corev1.ConditionFalse,
				LastUpdateTime: metav1.Now(),
				Reason:         "Cluster is not ready",
			})
		}
	} else {
		newSNConditions.RemoveCondition(v1alpha1.StorageNodeConditionTypeClusterReady)
	}

	// Update the available condition based on the phase
	if desiredState.Phase == v1alpha1.StorageNodePhaseReady {
		newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
			Type:           v1alpha1.StorageNodeConditionTypeAvailable,
			Status:         corev1.ConditionTrue,
			LastUpdateTime: metav1.Now(),
			Reason:         "All instances are ready",
		})
	} else {
		newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
			Type:           v1alpha1.StorageNodeConditionTypeAvailable,
			Status:         corev1.ConditionFalse,
			LastUpdateTime: metav1.Now(),
			Reason:         "One or more instances are not ready",
		})
	}

	// Update the registered condition
	if status.Registered {
		newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
			Type:           v1alpha1.StorageNodeConditionTypeRegistered,
			Status:         corev1.ConditionTrue,
			LastUpdateTime: metav1.Now(),
			Reason:         "StorageNode is registered",
		})
	} else {
		newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
			Type:           v1alpha1.StorageNodeConditionTypeRegistered,
			Status:         corev1.ConditionFalse,
			LastUpdateTime: metav1.Now(),
			Reason:         "StorageNode is not registered",
		})
	}
	return newSNConditions
}

// allInstancesReady returns true if all instances are ready, false otherwise
func allInstancesReady(instances []v1alpha1.InstanceStatus) bool {
	if len(instances) == 0 {
		return false
	}

	for idx := range instances {
		instance := &instances[idx]
		if !(instance.Status == string(rds.DBInstanceStatusAvailable)) {
			return false
		}
	}

	return true
}

func (r *StorageNodeReconciler) reconcileAwsRdsInstance(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	instance, err := client.GetInstance(ctx, node)
	if err != nil {
		return err
	}

	if instance == nil && node.Status.Phase != v1alpha1.StorageNodePhaseDeleting {
		err = client.CreateInstance(ctx, node, storageProvider.Spec.Parameters)
		if err != nil {
			return err
		}

		instance, err = client.GetInstance(ctx, node)
		if err != nil {
			return err
		}
	}

	if err := updateAWSRDSInstanceStatus(node, instance); err != nil {
		return fmt.Errorf("updateAWSRDSInstanceStatus failed: %w", err)
	}

	return nil
}

func updateAWSRDSInstanceStatus(node *v1alpha1.StorageNode, instance *rds.DescInstance) error {
	instances := make([]v1alpha1.InstanceStatus, 0)

	if instance == nil {
		node.Status.Instances = instances
		return nil
	}

	instances = append(instances, v1alpha1.InstanceStatus{
		Endpoint: v1alpha1.Endpoint{
			Address: instance.Endpoint.Address,
			Port:    instance.Endpoint.Port,
		},
		Status: string(instance.DBInstanceStatus),
	})

	node.Status.Instances = instances
	return nil
}

func (r *StorageNodeReconciler) reconcileAwsRDSCluster(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	ac, err := client.GetRDSCluster(ctx, node)
	if err != nil {
		return err
	}

	if ac == nil {
		// create instance
		err = client.CreateRDSCluster(ctx, node, storageProvider.Spec.Parameters)
		if err != nil {
			return err
		}
		ac, err = client.GetRDSCluster(ctx, node)
		if err != nil {
			return err
		}
	}

	// TODO: reconcile instance of aurora

	// update storage node status
	if err := updateClusterStatus(ctx, client, node, ac); err != nil {
		return fmt.Errorf("updateClusterStatus failed: %w", err)
	}

	return nil
}

func (r *StorageNodeReconciler) reconcileAwsAurora(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	r.Log.Info("reconcileAwsAurora", "node", node.GetName(), "phase", node.Status.Phase)
	ac, err := client.GetAuroraCluster(ctx, node)
	if err != nil {
		return err
	}

	if ac == nil {
		// create instance
		err = client.CreateAuroraCluster(ctx, node, storageProvider.Spec.Parameters)
		if err != nil {
			return err
		}
		ac, err = client.GetAuroraCluster(ctx, node)
		if err != nil {
			return err
		}
	}

	// TODO: reconcile instance of aurora

	// update storage node status
	if err := updateClusterStatus(ctx, client, node, ac); err != nil {
		return fmt.Errorf("updateClusterStatus failed: %w", err)
	}

	return nil
}

func updateClusterStatus(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, cluster *rds.DescCluster) error {
	// update cluster status
	clusterStatus := v1alpha1.ClusterStatus{}
	if cluster != nil {
		clusterStatus = v1alpha1.ClusterStatus{
			Status: cluster.Status,
			PrimaryEndpoint: v1alpha1.Endpoint{
				Address: cluster.PrimaryEndpoint,
				Port:    cluster.Port,
			},
			ReaderEndpoints: []v1alpha1.Endpoint{
				{
					Address: cluster.ReaderEndpoint,
					Port:    cluster.Port,
				},
			},
		}
	}
	node.Status.Cluster = clusterStatus

	// update instances status
	identifier := node.Annotations[v1alpha1.AnnotationsClusterIdentifier]
	filters := map[string][]string{
		"db-cluster-id": {identifier},
	}
	instances, err := client.GetInstancesByFilters(ctx, filters)
	if err != nil {
		return fmt.Errorf("GetInstances failed, err:%w", err)
	}

	var instanceStatus []v1alpha1.InstanceStatus
	for _, instance := range instances {
		instanceStatus = append(instanceStatus, v1alpha1.InstanceStatus{
			Status: string(instance.DBInstanceStatus),
			Endpoint: v1alpha1.Endpoint{
				Address: instance.Endpoint.Address,
				Port:    instance.Endpoint.Port,
			}})
	}
	node.Status.Instances = instanceStatus
	return nil
}

// deleteDatabaseCluster
func (r *StorageNodeReconciler) deleteDatabaseCluster(ctx context.Context, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	switch storageProvider.Spec.Provisioner {
	case v1alpha1.ProvisionerAWSRDSInstance:
		if err := r.deleteAWSRDSInstance(ctx, aws.NewRdsClient(r.AwsRDS), node, storageProvider); err != nil {
			return fmt.Errorf("delete aws rds instance failed: %w", err)
		}
	case v1alpha1.ProvisionerAWSRDSCluster:
		if err := r.deleteAWSRDSCluster(ctx, aws.NewRdsClient(r.AwsRDS), node, storageProvider); err != nil {
			return fmt.Errorf("delete aws rds cluster failed: %w", err)
		}
	case v1alpha1.ProvisionerAWSAurora:
		if err := r.deleteAWSAurora(ctx, aws.NewRdsClient(r.AwsRDS), node, storageProvider); err != nil {
			return fmt.Errorf("delete aws aurora cluster failed: %w", err)
		}
	default:
		return fmt.Errorf("unsupported database provisioner %s", storageProvider.Spec.Provisioner)
	}
	return nil
}

func (r *StorageNodeReconciler) deleteAWSRDSInstance(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	if node.Annotations[v1alpha1.AnnotationsInstanceIdentifier] == "" {
		return nil
	}

	instance, err := client.GetInstance(ctx, node)
	if err != nil {
		return err
	}

	if instance != nil && instance.DBInstanceStatus != rds.DBInstanceStatusDeleting {
		if err := client.DeleteInstance(ctx, node, storageProvider); err != nil {
			r.Recorder.Eventf(node, corev1.EventTypeWarning, "DeleteFailed", "Failed to delete instance %s: %s", node.Annotations[v1alpha1.AnnotationsInstanceIdentifier], err.Error())
			return err
		}
		r.Recorder.Event(node, corev1.EventTypeNormal, "Deleting", fmt.Sprintf("instance %s is deleting", node.Annotations[v1alpha1.AnnotationsInstanceIdentifier]))
	}

	// update instance status
	if err := updateAWSRDSInstanceStatus(node, instance); err != nil {
		return fmt.Errorf("updateAWSRDSInstanceStatus failed: %w", err)
	}

	return nil
}

// nolint:dupl
func (r *StorageNodeReconciler) deleteAWSRDSCluster(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	if node.Annotations[v1alpha1.AnnotationsClusterIdentifier] == "" {
		return nil
	}

	cluster, err := client.GetRDSCluster(ctx, node)
	if err != nil {
		return fmt.Errorf("get rds cluster failed: %w", err)
	}
	if cluster != nil && cluster.Status != string(rds.DBClusterStatusDeleting) {
		if err := client.DeleteRDSCluster(ctx, node, storageProvider); err != nil {
			r.Recorder.Eventf(node, corev1.EventTypeWarning, "DeleteFailed", "Failed to delete rds cluster %s: %s", node.Annotations[v1alpha1.AnnotationsClusterIdentifier], err.Error())
			return err
		}
		r.Recorder.Event(node, corev1.EventTypeNormal, "Deleting", fmt.Sprintf("rds cluster %s is deleting", node.Annotations[v1alpha1.AnnotationsClusterIdentifier]))
	}

	// update storage node status
	if err := updateClusterStatus(ctx, client, node, cluster); err != nil {
		return fmt.Errorf("updateClusterStatus failed: %w", err)
	}
	return nil
}

// nolint:dupl
func (r *StorageNodeReconciler) deleteAWSAurora(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	if node.Annotations[v1alpha1.AnnotationsClusterIdentifier] == "" {
		return nil
	}

	auroraCluster, err := client.GetAuroraCluster(ctx, node)
	if err != nil {
		return fmt.Errorf("get aurora cluster failed: %w", err)
	}
	if auroraCluster != nil && auroraCluster.Status != string(rds.DBClusterStatusDeleting) {
		if err := client.DeleteAuroraCluster(ctx, node, storageProvider); err != nil {
			r.Recorder.Eventf(node, corev1.EventTypeWarning, "DeleteFailed", "Failed to delete aurora cluster %s: %s", node.Annotations[v1alpha1.AnnotationsClusterIdentifier], err.Error())
			return err
		}
		r.Recorder.Event(node, corev1.EventTypeNormal, "Deleting", fmt.Sprintf("aurora cluster %s is deleting", node.Annotations[v1alpha1.AnnotationsClusterIdentifier]))
	}

	// update storage node status
	if err := updateClusterStatus(ctx, client, node, auroraCluster); err != nil {
		return fmt.Errorf("updateClusterStatus failed: %w", err)
	}
	return nil
}

// registerStorageUnit
func (r *StorageNodeReconciler) registerStorageUnit(ctx context.Context, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	// if register storage unit is not enabled, return
	if node.Annotations[AnnotationKeyRegisterStorageUnitEnabled] != "true" {
		return nil
	}

	if err := r.validateComputeNodeAnnotations(node); err != nil {
		return err
	}

	// if storage unit is already registered, return
	if node.Status.Registered {
		return nil
	}

	// if node is not ready, return
	if node.Status.Phase != v1alpha1.StorageNodePhaseReady {
		r.Recorder.Eventf(node, corev1.EventTypeWarning, "RegisterWaiting", "Waiting to register storage unit for node %s/%s: node is not ready", node.GetNamespace(), node.GetName())
		return nil
	}

	logicDBName := node.Annotations[AnnotationKeyLogicDatabaseName]
	dbName := node.Annotations[v1alpha1.AnnotationsInstanceDBName]

	ssServer, err := r.getShardingsphereServer(ctx, node)
	if err != nil {
		return fmt.Errorf("getShardingsphereServer failed: %w", err)
	}

	defer ssServer.Close()

	if err := ssServer.CreateDatabase(logicDBName); err != nil {
		return fmt.Errorf("create database failed: %w", err)
	}
	r.Recorder.Eventf(node, corev1.EventTypeNormal, "LogicDatabaseCreated", "LogicDatabase %s is created", logicDBName)

	var host string
	var port int32
	var username, password string
	// get storage unit info from instance
	if node.Status.Cluster.Status == "" {
		host, port, username, password = getDatasourceInfoFromInstance(node, storageProvider)
	} else {
		host, port, username, password = getDatasourceInfoFromCluster(node, storageProvider)
	}

	if err := ssServer.RegisterStorageUnit(logicDBName, getDSName(node), host, uint(port), dbName, username, password); err != nil {
		return fmt.Errorf("register storage node failed: %w", err)
	}
	r.Recorder.Eventf(node, corev1.EventTypeNormal, "StorageUnitRegistered", "StorageUnit %s:%d/%s is registered", host, port, dbName)

	node.Status.Registered = true
	return nil
}

// getDSName returns the datasource name of the storage node.
// datasource name only allows letters, numbers and _, and must start with a letter.
// ref: https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-proxy/distsql/syntax/rdl/storage-unit-definition/register-storage-unit/
func getDSName(node *v1alpha1.StorageNode) string {
	return fmt.Sprintf("ds_%s", strings.ReplaceAll(node.GetName(), "-", "_"))
}

func getDatasourceInfoFromInstance(node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) (host string, port int32, username, password string) {
	ins := node.Status.Instances[0]
	host = ins.Endpoint.Address
	port = ins.Endpoint.Port
	username = node.Annotations[v1alpha1.AnnotationsMasterUsername]
	if username == "" {
		username = storageProvider.Spec.Parameters["masterUsername"]
	}
	password = node.Annotations[v1alpha1.AnnotationsMasterUserPassword]
	if password == "" {
		password = storageProvider.Spec.Parameters["masterUserPassword"]
	}
	return
}

func getDatasourceInfoFromCluster(node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) (host string, port int32, username, password string) {
	cluster := node.Status.Cluster
	host = cluster.PrimaryEndpoint.Address
	port = cluster.PrimaryEndpoint.Port
	username = node.Annotations[v1alpha1.AnnotationsMasterUsername]
	if username == "" {
		username = storageProvider.Spec.Parameters["masterUsername"]
	}
	password = node.Annotations[v1alpha1.AnnotationsMasterUserPassword]
	if password == "" {
		password = storageProvider.Spec.Parameters["masterUserPassword"]
	}
	return
}

func (r *StorageNodeReconciler) unregisterStorageUnit(ctx context.Context, node *v1alpha1.StorageNode) error {
	if !node.Status.Registered {
		return nil
	}
	if err := r.validateComputeNodeAnnotations(node); err != nil {
		return err
	}

	logicDBName := node.Annotations[AnnotationKeyLogicDatabaseName]

	ssServer, err := r.getShardingsphereServer(ctx, node)
	if err != nil {
		return fmt.Errorf("getShardingsphereServer failed: %w", err)
	}

	defer ssServer.Close()

	if err := ssServer.UnRegisterStorageUnit(logicDBName, getDSName(node)); err != nil {
		return fmt.Errorf("unregister storage unit failed: %w", err)
	}

	r.Recorder.Eventf(node, corev1.EventTypeNormal, "StorageUnitUnRegistered", "StorageUnit of node %s/%s is unregistered", node.GetNamespace(), node.GetName())

	node.Status.Registered = false
	return nil
}

func (r *StorageNodeReconciler) validateComputeNodeAnnotations(node *v1alpha1.StorageNode) error {
	requiredAnnos := []string{
		AnnotationKeyLogicDatabaseName,
		v1alpha1.AnnotationsInstanceDBName,
		AnnotationKeyComputeNodeName,
	}

	for _, anno := range requiredAnnos {
		if v, ok := node.Annotations[anno]; !ok || v == "" {
			r.Recorder.Eventf(node, corev1.EventTypeWarning, "RegisterChecking", "Waiting to register storage unit for node %s/%s: annotation %s is required", node.GetNamespace(), node.GetName(), anno)
			return fmt.Errorf("annotation %s is required", anno)
		}
	}

	return nil
}

func (r *StorageNodeReconciler) getShardingsphereServer(ctx context.Context, node *v1alpha1.StorageNode) (shardingsphere.IServer, error) {
	var (
		driver, host, username, password string
		port                             uint
	)

	// get compute node
	cn := &v1alpha1.ComputeNode{}
	if err := r.Client.Get(ctx, types.NamespacedName{
		Name:      node.Annotations[AnnotationKeyComputeNodeName],
		Namespace: node.Namespace,
	}, cn); err != nil {
		return nil, fmt.Errorf("get compute node failed: %w", err)
	}

	serverConf := cn.Spec.Bootstrap.ServerConfig

	driver, ok := serverConf.Props[ShardingSphereProtocolType]
	if !ok || driver == "" {
		driver = "mysql"
	}
	driver = strings.ToLower(driver)

	if len(serverConf.Authority.Users) == 0 {
		return nil, fmt.Errorf("no user in compute node %s/%s", cn.Namespace, cn.Name)
	}

	username = strings.Split(serverConf.Authority.Users[0].User, "@")[0]
	password = serverConf.Authority.Users[0].Password

	// get service of compute node
	svc, err := r.Service.GetByNamespacedName(ctx, types.NamespacedName{
		Name:      node.Annotations[AnnotationKeyComputeNodeName],
		Namespace: node.Namespace,
	})

	if err != nil || svc == nil {
		return nil, fmt.Errorf("get service failed: %w", err)
	}

	host = fmt.Sprintf("%s.%s", svc.Name, svc.Namespace)

	port = uint(svc.Spec.Ports[0].Port)

	ssServer, err := shardingsphere.NewServer(driver, host, port, username, password)
	if err != nil {
		return nil, fmt.Errorf("new shardingsphere server failed: %w", err)
	}

	return ssServer, nil
}

func (r *StorageNodeReconciler) reconcileCloudNativePG(ctx context.Context, sn *v1alpha1.StorageNode, sp *v1alpha1.StorageProvider) error {
	cluster, err := r.getCloudNativePGCluster(ctx, types.NamespacedName{Namespace: sn.Namespace, Name: sn.Name})
	if err != nil {
		return err
	}
	if cluster != nil {
		return r.updateCloudNativePGCluster(ctx, sn, sp, cluster)
	}
	return r.createCloudNativePGCluster(ctx, sn, sp)
}

func (r *StorageNodeReconciler) getCloudNativePGCluster(ctx context.Context, namespacedName types.NamespacedName) (*cnpg.Cluster, error) {
	c, err := r.CNPG.GetClusterByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *StorageNodeReconciler) createCloudNativePGCluster(ctx context.Context, sn *v1alpha1.StorageNode, sp *v1alpha1.StorageProvider) error {
	cluster := r.CNPG.Build(ctx, sn, sp)
	err := r.CNPG.Create(ctx, cluster)
	if err != nil && apierrors.IsAlreadyExists(err) || err == nil {
		return nil
	}
	return err
}

func (r *StorageNodeReconciler) updateCloudNativePGCluster(ctx context.Context, sn *v1alpha1.StorageNode, sp *v1alpha1.StorageProvider, cluster *cnpg.Cluster) error {
	exp := r.CNPG.Build(ctx, sn, sp)
	exp.ObjectMeta = cluster.ObjectMeta
	exp.Labels = cluster.Labels
	exp.Annotations = cluster.Annotations

	if !reflect.DeepEqual(cluster.Spec, exp.Spec) {
		return r.CNPG.Update(ctx, exp)
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager
func (r *StorageNodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.StorageNode{}).
		Complete(r)
}
