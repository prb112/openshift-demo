/*
Copyright IBM Corp. 2023

SPDX-License-Identifier: Apache-2.0
*/

package controller

import (
	"context"
	batch "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	externalmonitorocppowerxyzv1alpha1 "github.com/prb112/openshift-demo/operator/external-monitor/api/v1alpha1"
)

// MonitorReconciler reconciles a Monitor object
type MonitorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=external-monitor.ocp-power.xyz,resources=monitors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=external-monitor.ocp-power.xyz,resources=monitors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=external-monitor.ocp-power.xyz,resources=monitors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims the controller
func (r *MonitorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	// Fetch the Monitor instance
	monitor := &externalmonitorocppowerxyzv1alpha1.Monitor{}
	err := r.Get(ctx, req.NamespacedName, monitor)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			l.Info("Monitor resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		l.Error(err, "Failed to get Monitor")
		return ctrl.Result{}, err
	} else {
		l.WithValues("key", monitor).V(0).Info("XReached this part of the code")
	}

	// Check if the CronJob already exists, if not create a new one
	found := &batch.CronJob{}
	l.V(0).Info("Reached this part of the code")
	_ = found
	/*
		apiVersion: batch/v1
		kind: CronJob
		metadata:
		  name: hello
		spec:
		  schedule: "* * * * *"
		  jobTemplate:
		    spec:
		      template:
		        spec:
		          containers:
		          - name: hello
		            image: busybox:1.28
		            imagePullPolicy: IfNotPresent
		            command:
		            - /bin/sh
		            - -c
		            - date; echo Hello from the Kubernetes cluster
		          restartPolicy: OnFailure
	*/
	//monitor.Spec.Tag
	//	monitor.Spec.Name

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&externalmonitorocppowerxyzv1alpha1.Monitor{}).
		Complete(r)
}
