/*
Copyright IBM Corp. 2023

SPDX-License-Identifier: Apache-2.0
*/

package controller

import (
	"context"
	"fmt"
	"github.com/prb112/openshift-demo/operator/external-monitor/bindata"
	v1 "k8s.io/api/batch/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"strconv"

	"strings"
	"time"

	externalmonitorocppowerxyzv1alpha1 "github.com/prb112/openshift-demo/operator/external-monitor/api/v1alpha1"
)

// MonitorReconciler reconciles a Monitor object
type MonitorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

var (
	appsScheme = runtime.NewScheme()
	appsCodecs = serializer.NewCodecFactory(appsScheme)
)

//+kubebuilder:rbac:groups=external-monitor.ocp-power.xyz,resources=monitors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=external-monitor.ocp-power.xyz,resources=monitors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=external-monitor.ocp-power.xyz,resources=monitors/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=cronjobs,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=job,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims the controller
func (r *MonitorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.V(0).Info("Start - 012345")
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

	// Function
	go func() {
		seconds, _ := time.ParseDuration("30s")
		for {
			opts := []client.ListOption{
				client.InNamespace("default"),
			}
			list := &v1.CronJobList{}
			err2 := r.List(ctx, list, opts...)
			if err2 != nil {
				l.V(0).Info("Checking the CronJob for is-up-nfs status - not configured")
				l.Error(err2, "ugh")
			} else {
				prefix := ""
				for _, item := range list.Items {
					l.WithValues("key", item).V(0).Info("Here is item")
					for _, v := range item.Status.Active {
						l.WithValues("key", v.Name).V(0).Info("Here is the Job Name")
						prefix = v.Name
					}
				}

				// prefix is used to select the Pod
				opts := []client.ListOption{
					client.InNamespace("default"),
				}
				list := &v12.PodList{}
				err32 := r.List(ctx, list, opts...)
				if err32 != nil {
					l.V(0).Info("Checking the Pod List for is-up-nfs status - not configured")
					l.Error(err32, "ugh")
				} else {
					var lastFailed bool
					for _, pod := range list.Items {
						if strings.HasPrefix(pod.Name, prefix) {
							l.WithValues("pod", pod).V(0).Info("Found prefix")
							if len(pod.Status.ContainerStatuses) != 0 &&
								pod.Status.ContainerStatuses[0].State.Terminated != nil &&
								pod.Status.ContainerStatuses[0].State.Terminated.ExitCode != 0 {
								lastFailed = true
							} else {
								lastFailed = false
							}
							l.WithValues("failed", lastFailed).V(0).Info("here?")
						}
					}
					if lastFailed {

						// job.yaml
						b, _ := bindata.Asset("assets/job.yaml")
						decoder := serializer.NewCodecFactory(scheme.Scheme).UniversalDecoder()
						recoveryJob := &v1.Job{}

						err := runtime.DecodeInto(decoder, b, recoveryJob)
						if err != nil {
							fmt.Println(err)
						}
						//recoveryJob.Namespace = "default"
						recoveryJob.Name = recoveryJob.Name + "-" + strconv.Itoa(int(time.Now().UnixMilli()))
						//recoveryJob.Spec.Template.Spec.RestartPolicy = "Never"

						errx := r.Create(ctx, recoveryJob)
						l.WithValues("recoveryJob", recoveryJob).V(0).Info("recovery job created?")
						l.V(0).Info("Successful creation")
						fmt.Println(errx)

					}
				}

				l.V(0).Info("Checking the CronJob for is-up-nfs status")
			}

			l.V(0).Info("Checking the CronJob")

			time.Sleep(seconds)
		}
	}()

	// Check if the CronJob already exists, if not create a new one
	//found := &batch.CronJob{}
	//l.V(0).Info("Reached this part of the code")

	return ctrl.Result{}, nil
}

func (r *MonitorReconciler) findObjectsForCronJob(cronJob client.Object) []reconcile.Request {
	cronJobs := &v1.CronJobList{}
	listOps := &client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", cronJob.GetName()),
		Namespace:     cronJob.GetNamespace(),
	}
	err := r.List(context.TODO(), cronJobs, listOps)
	if err != nil {
		return []reconcile.Request{}
	}

	requests := make([]reconcile.Request, len(cronJobs.Items))
	for i, item := range cronJobs.Items {
		requests[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      item.GetName(),
				Namespace: item.GetNamespace(),
			},
		}
	}
	return requests
}

// SetupWithManager sets up the controller with the Manager.
func (r *MonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		Watches(
			&source.Kind{Type: &v1.CronJob{}},
			handler.EnqueueRequestsFromMapFunc(r.findObjectsForCronJob),
		).
		Owns(&v1.CronJob{}).
		For(&externalmonitorocppowerxyzv1alpha1.Monitor{}).
		Complete(r)
}
