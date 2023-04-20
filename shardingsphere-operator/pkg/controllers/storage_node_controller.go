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

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/aws"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	dbmeshv1alpha1 "github.com/database-mesh/golang-sdk/kubernetes/api/v1alpha1"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	StorageNodeControllerName = "storage-node-controller"
	FinalizerName             = "shardingsphere.apache.org/finalizer"
)

// StorageNodeReconciler is a controller for storage nodes
type StorageNodeReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Log      logr.Logger
	Recorder record.EventRecorder
	AwsRDS   rds.RDS
}

// Reconcile handles main function of this controller
// nolint:gocognit
func (r *StorageNodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(StorageNodeControllerName, req.NamespacedName)

	logger.Info("Reconciling StorageNode")

	node := &v1alpha1.StorageNode{}
	if err := r.Get(ctx, req.NamespacedName, node); err != nil {
		if client.IgnoreNotFound(err) == nil {
			logger.Info(fmt.Sprintf("StorageNode [%s:%s] is not exist", req.Namespace, req.Name))
			return ctrl.Result{}, nil
		}
		logger.Error(err, fmt.Sprintf("unable to fetch StorageNode [%s:%s]", req.Namespace, req.Name))
		return ctrl.Result{Requeue: true}, err
	}

	// Get databaseClass with storageNode.Spec.DatabaseClassName
	databaseClass, err := r.getDatabaseClass(ctx, node)
	if err != nil {
		logger.Error(err, fmt.Sprintf("unable to fetch DatabaseClass [%s]", node.Spec.DatabaseClassName))
		return ctrl.Result{Requeue: true}, err
	}

	// TODO: when storage node needed finalized, set deletion timestamp and set status to deleting, waiting database instance deleted. and then remove finalizer and delete storage node.

	// finalize storage node
	if err := r.finalize(ctx, node, databaseClass); err != nil {
		logger.Error(err, fmt.Sprintf("unable to finalize StorageNode [%s:%s]", node.GetNamespace(), node.GetName()))
		return ctrl.Result{}, err
	}

	// reconcile storage node with databaseClass
	switch databaseClass.Spec.Provisioner {
	case dbmeshv1alpha1.ProvisionerAWSRDSInstance:
		if err := r.reconcileAwsRdsInstance(ctx, aws.NewRdsClient(r.AwsRDS), node, databaseClass); err != nil {
			logger.Error(err, fmt.Sprintf("unable to reconcile AWS RDS Instance [%s:%s], err:%s", node.GetNamespace(), node.GetName(), err.Error()))
			r.Recorder.Event(node, corev1.EventTypeWarning, fmt.Sprintf("Reconcile [%s:%s] Failed", node.GetNamespace(), node.GetName()), err.Error())
		}
	case dbmeshv1alpha1.ProvisionerAWSAurora:
		if err := r.reconcileAwsAurora(ctx, aws.NewRdsClient(r.AwsRDS), node, databaseClass); err != nil {
			r.Recorder.Event(node, corev1.EventTypeWarning, fmt.Sprintf("Reconcile [%s:%s] Failed", node.GetNamespace(), node.GetName()), err.Error())
		}
	default:
		r.Recorder.Event(node, corev1.EventTypeWarning, "UnsupportedDatabaseProvisioner", fmt.Sprintf("unsupported database provisioner [%s]", databaseClass.Spec.Provisioner))
		logger.Error(nil, fmt.Sprintf("unsupported database provisioner [%s]", databaseClass.Spec.Provisioner))
	}

	// update status
	desiredState := computeDesiredState(node.Status)

	if !reflect.DeepEqual(node.Status, desiredState) {
		node.Status = desiredState
		err := r.Status().Update(ctx, node)
		if err != nil {
			logger.Error(err, fmt.Sprintf("unable to update StorageNode [%s:%s] status", req.Namespace, req.Name))
			return ctrl.Result{Requeue: true}, err
		}
	}

	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
}

func (r *StorageNodeReconciler) getDatabaseClass(ctx context.Context, node *v1alpha1.StorageNode) (databaseClass *dbmeshv1alpha1.DatabaseClass, err error) {
	databaseClass = &dbmeshv1alpha1.DatabaseClass{}

	if err := r.Get(ctx, client.ObjectKey{Name: node.Spec.DatabaseClassName}, databaseClass); err != nil {
		r.Log.Error(err, fmt.Sprintf("unable to fetch DatabaseClass [%s]", node.Spec.DatabaseClassName))
		r.Recorder.Event(node, corev1.EventTypeWarning, "DatabaseClassNotFound", fmt.Sprintf("DatabaseClass [%s] not found", node.Spec.DatabaseClassName))
		return nil, err
	}

	// check provisioner
	// aws-like provisioner need aws rds client
	if databaseClass.Spec.Provisioner == dbmeshv1alpha1.ProvisionerAWSRDSInstance || databaseClass.Spec.Provisioner == dbmeshv1alpha1.ProvisionerAWSAurora {
		if r.AwsRDS == nil {
			r.Recorder.Event(node, corev1.EventTypeWarning, "AwsRdsClientIsNil", "aws rds client is nil, please check your aws credentials")
			return nil, fmt.Errorf("aws rds client is nil, please check your aws credentials")
		}
	}

	return databaseClass, nil
}

func removeString(finalizers []string, name string) []string {
	var result []string
	for _, finalizer := range finalizers {
		if finalizer != name {
			result = append(result, finalizer)
		}
	}
	return result
}

func containsString(finalizers []string, name string) bool {
	for _, finalizer := range finalizers {
		if finalizer == name {
			return true
		}
	}
	return false
}

// nolint:nestif
func (r *StorageNodeReconciler) finalize(ctx context.Context, node *v1alpha1.StorageNode, databaseClass *dbmeshv1alpha1.DatabaseClass) error {
	if node.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent to registering our finalizer.
		if !containsString(node.ObjectMeta.Finalizers, FinalizerName) {
			node.ObjectMeta.Finalizers = append(node.ObjectMeta.Finalizers, FinalizerName)
			if err := r.Update(ctx, node); err != nil {
				return err
			}
		}
	} else if containsString(node.ObjectMeta.Finalizers, FinalizerName) {
		// The object is being deleted
		if err := r.deleteDatabaseCluster(ctx, node, databaseClass); err != nil {
			return err
		}
		// remove our finalizer from the list and update it.
		node.ObjectMeta.Finalizers = removeString(node.ObjectMeta.Finalizers, FinalizerName)
		if err := r.Update(ctx, node); err != nil {
			return err
		}
	}
	return nil
}

// nolint:gocritic
func computeDesiredState(status v1alpha1.StorageNodeStatus) v1alpha1.StorageNodeStatus {
	// Initialize a new status object based on the current state
	desiredState := status

	// TODO: set enums for aws instance status

	// If the cluster status is not empty, then we compute the phase based on the cluster status
	clusterStatus := ""
	if status.Cluster.Status != "" {
		if status.Cluster.Status == "available" {
			clusterStatus = "Ready"
		}
	}

	if (clusterStatus == "" || clusterStatus == "Ready") && allInstancesReady(status.Instances) {
		desiredState.Phase = v1alpha1.StorageNodePhaseReady
	} else {
		desiredState.Phase = v1alpha1.StorageNodePhaseNotReady
	}

	newSNConditions := status.Conditions

	// Update the cluster ready condition if the cluster status is not empty
	if clusterStatus != "" {
		if clusterStatus == "Ready" {
			newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
				Type:   v1alpha1.StorageNodeConditionTypeClusterReady,
				Status: corev1.ConditionTrue,
				Reason: "Cluster is ready",
			})
		} else {
			newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
				Type:   v1alpha1.StorageNodeConditionTypeClusterReady,
				Status: corev1.ConditionFalse,
				Reason: "Cluster is not ready",
			})
		}
	} else {
		newSNConditions.RemoveCondition(v1alpha1.StorageNodeConditionTypeClusterReady)
	}

	// Update the available condition based on the phase
	if desiredState.Phase == v1alpha1.StorageNodePhaseReady {
		newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
			Type:   v1alpha1.StorageNodeConditionTypeAvailable,
			Status: corev1.ConditionTrue,
			Reason: "All instances are ready",
		})
	} else {
		newSNConditions.UpsertCondition(&v1alpha1.StorageNodeCondition{
			Type:   v1alpha1.StorageNodeConditionTypeAvailable,
			Status: corev1.ConditionFalse,
			Reason: "One or more instances are not ready",
		})
	}

	desiredState.Conditions = newSNConditions

	return desiredState
}

// allInstancesReady returns true if all instances are ready, false otherwise
func allInstancesReady(instances []v1alpha1.InstanceStatus) bool {
	if len(instances) == 0 {
		return false
	}

	for idx := range instances {
		instance := &instances[idx]
		if !(instance.Status == "Ready") {
			return false
		}
	}

	return true
}

func (r *StorageNodeReconciler) reconcileAwsRdsInstance(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, dbClass *dbmeshv1alpha1.DatabaseClass) error {
	instance, err := client.GetInstance(ctx, node)
	if err != nil {
		return err
	}

	if instance == nil {
		err = client.CreateInstance(ctx, node, dbClass.Spec.Parameters)
		if err != nil {
			return err
		}

		instance, err = client.GetInstance(ctx, node)
		if err != nil {
			return err
		}
	}

	r.Log.Info(fmt.Sprintf("RDS instance [%s] status is [%s]", instance.DBInstanceIdentifier, instance.DBInstanceStatus))

	newStatus := updateInstanceStatus(node, instance)
	node.Status.Instances = newStatus
	if err := r.Status().Update(ctx, node); err != nil {
		r.Log.Error(err, fmt.Sprintf("Failed to update status for node [%s:%s]", node.GetNamespace(), node.GetName()))
	}
	r.Recorder.Eventf(node, corev1.EventTypeNormal, "Reconcile", "Reconciled RDS instance %s, status is %s", instance.DBInstanceIdentifier, instance.DBInstanceStatus)
	return nil
}

func updateInstanceStatus(node *v1alpha1.StorageNode, instance *rds.DescInstance) []v1alpha1.InstanceStatus {
	instances := make([]v1alpha1.InstanceStatus, 0)

	status := instance.DBInstanceStatus
	if status == "available" {
		status = "Ready"
	}

	instances = append(instances, v1alpha1.InstanceStatus{
		Endpoint: v1alpha1.Endpoint{
			Address: instance.Endpoint.Address,
			Port:    instance.Endpoint.Port,
		},
		Status: status,
	})
	return instances
}

func (r *StorageNodeReconciler) reconcileAwsAurora(ctx context.Context, client aws.IRdsClient, node *v1alpha1.StorageNode, dbClass *dbmeshv1alpha1.DatabaseClass) error {
	// get instance
	instance, err := client.GetAuroraCluster(ctx, node)
	if err != nil {
		return err
	}
	if instance == nil {
		// create instance
		err = client.CreateAuroraCluster(ctx, node, dbClass.Spec.Parameters)
		if err != nil {
			return err
		}
	}
	// TODO: update storage node status
	return nil
}

// deleteDatabaseCluster
func (r *StorageNodeReconciler) deleteDatabaseCluster(ctx context.Context, node *v1alpha1.StorageNode, databaseClass *dbmeshv1alpha1.DatabaseClass) error {
	switch databaseClass.Spec.Provisioner {
	case dbmeshv1alpha1.ProvisionerAWSRDSInstance:
		if err := aws.NewRdsClient(r.AwsRDS).DeleteInstance(ctx, node, databaseClass); err != nil {
			return err
		}
	case dbmeshv1alpha1.ProvisionerAWSAurora:
		if err := aws.NewRdsClient(r.AwsRDS).DeleteAuroraCluster(ctx, node, databaseClass); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported database provisioner [%s]", databaseClass.Spec.Provisioner)
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager
func (r *StorageNodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.StorageNode{}).
		Complete(r)
}
