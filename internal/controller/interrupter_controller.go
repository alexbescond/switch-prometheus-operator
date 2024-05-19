/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	switchv1beta1 "github.com/alexbescond/switch-prometheus-operator/api/v1beta1"
)

// InterrupterReconciler reconciles a Interrupter object
type InterrupterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=switch.alexbescond.io,resources=interrupters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=switch.alexbescond.io,resources=interrupters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=switch.alexbescond.io,resources=interrupters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Interrupter object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *InterrupterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	var interrupter switchv1beta1.Interrupter
	if err := r.Get(ctx, req.NamespacedName, &interrupter); err != nil {
		log.Log.Error(err, "unable to fetch Interrupter")
		return ctrl.Result{}, err
	}

	log.Log.Info("Reconciling Interrupter", "Interrupter", interrupter)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *InterrupterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&switchv1beta1.Interrupter{}).
		Complete(r)
}
