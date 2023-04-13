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
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/awsaurora"
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

	node := &v1alpha1.StorageNode{}
	if err := r.Get(ctx, req.NamespacedName, node); err != nil {
		if client.IgnoreNotFound(err) != nil {
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
		return ctrl.Result{Requeue: true}, err
	}

	desiredState := computeDesiredState(node.Status)

	if !reflect.DeepEqual(node.Status, desiredState) {
		node.Status = desiredState
		err := r.Status().Update(ctx, node)
		if err != nil {
			return ctrl.Result{Requeue: true}, err
		}
	}

	// Get databaseClass with storageNode.Spec.DatabaseClassName
	// if databaseClass can not be found, requeue and set a warning event to storageNode
	databaseClass := &dbmeshv1alpha1.DatabaseClass{}
	if err := r.Get(ctx, client.ObjectKey{Name: node.Spec.DatabaseClassName}, databaseClass); err != nil {
		logger.Error(err, fmt.Sprintf("unable to fetch DatabaseClass [%s]", node.Spec.DatabaseClassName))
		r.Recorder.Event(node, corev1.EventTypeWarning, "DatabaseClassNotFound", fmt.Sprintf("DatabaseClass [%s] not found", node.Spec.DatabaseClassName))
		return ctrl.Result{RequeueAfter: defaultRequeueTime}, err
	}

	var cClient storagenode.IDBClusterClient
	switch databaseClass.Spec.Provisioner {
	case "aws-aurora":
		cClient = awsaurora.New(r.AwsRDS)
	default:
		logger.Error(nil, fmt.Sprintf("unsupported database provisioner [%s]", databaseClass.Spec.Provisioner))
	}

	if err := cClient.IsValid(node); err != nil {
		logger.Error(err, fmt.Sprintf("invalid database cluster [%s]", node.Spec.DatabaseClassName))
		return ctrl.Result{Requeue: true}, err
	}

	// finalize
	// nolint:nestif
	if node.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent to registering our finalizer.
		if !containsString(node.ObjectMeta.Finalizers, FinalizerName) {
			node.ObjectMeta.Finalizers = append(node.ObjectMeta.Finalizers, FinalizerName)
			err := r.Update(ctx, node)
			if err != nil {
				return ctrl.Result{Requeue: true}, err
			}
		}
	} else if containsString(node.ObjectMeta.Finalizers, FinalizerName) {
		if err := r.deleteDatabaseCluster(ctx, node, cClient); err != nil {
			return ctrl.Result{Requeue: true}, err
		}
		node.ObjectMeta.Finalizers = removeString(node.ObjectMeta.Finalizers, FinalizerName)
		err := r.Update(ctx, node)
		if err != nil {
			return ctrl.Result{Requeue: true}, err
		}
	}

	// reconcile database cluster
	if err := r.reconcileDatabaseCluster(ctx, node, cClient, databaseClass.Spec.Parameters); err != nil {
		logger.Error(err, fmt.Sprintf("unable to reconcile DatabaseCluster [%s], err:%s", node.Spec.DatabaseClassName, err.Error()))
		return ctrl.Result{Requeue: true}, err
	}

	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
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

// nolint
func computeDesiredState(status v1alpha1.StorageNodeStatus) v1alpha1.StorageNodeStatus {
	// Initialize a new status object based on the current state
	desiredState := status

	// Compute the desired phase based on the number of instances and their readiness
	if status.Cluster.Status == "Ready" && allInstancesReady(status.Instances) {
		desiredState.Phase = v1alpha1.StorageNodePhaseReady
		desiredState.Cluster.Status = "Ready"
	} else {
		desiredState.Phase = v1alpha1.StorageNodePhaseNotReady
		desiredState.Cluster.Status = "NotReady"
	}

	// Compute the desired conditions based on the phase and any errors
	newSNConditions := status.Conditions

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
	// nolint
	for _, instance := range instances {
		if !(instance.Status == "Ready") {
			return false
		}
	}
	return true
}

func (r *StorageNodeReconciler) reconcileDatabaseCluster(ctx context.Context, node *v1alpha1.StorageNode, client storagenode.IDBClusterClient, params map[string]string) error {
	cluster, err := client.GetCluster(ctx, node)
	if err != nil {
		r.Log.Error(err, fmt.Sprintf("unable to get database cluster [%s]", node.Spec.DatabaseClassName))
		return err
	}

	if cluster == nil {
		cluster, err = client.CreateCluster(ctx, node, params)
		if err != nil {
			r.Log.Error(err, fmt.Sprintf("unable to create database cluster [%s]", node.Spec.DatabaseClassName))
			return err
		}
	}
	r.Log.Info(fmt.Sprintf("database cluster [%+v] is ready", cluster))
	return nil
}

// deleteDatabaseCluster when a storageNode is deleted, we need to delete the database cluster
func (r *StorageNodeReconciler) deleteDatabaseCluster(ctx context.Context, node *v1alpha1.StorageNode, client storagenode.IDBClusterClient) error {
	cluster, err := client.GetCluster(ctx, node)
	if err != nil {
		return err
	}
	if cluster == nil {
		return nil
	}
	// delete database cluster
	if err := client.DeleteCluster(ctx, node); err != nil {
		return err
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager
func (r *StorageNodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.StorageNode{}).
		Complete(r)
}
