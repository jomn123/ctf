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
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctfv1alpha1 "securinetes.com/ctf/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// CloudReconciler reconciles a Cloud object
type CloudReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ctf.securinetes.com,resources=clouds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ctf.securinetes.com,resources=clouds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ctf.securinetes.com,resources=clouds/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Cloud object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *CloudReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	ctf := ctfv1alpha1.Cloud{}
	if err := r.Get(ctx, req.NamespacedName, &ctf); err != nil {
		log.Error(err, "you are not qualified yet...")

		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if ctf.Spec.Intercept == "true" {
		fmt.Printf("You intercepted a message from the Rikage to the Mizukage ...\n")
		time.Sleep(30 * time.Second)
		fmt.Printf("They will meet in a secret location ..\n")
		if err := r.createriddle(ctx); err != nil {
			return ctrl.Result{}, err
		}
		time.Sleep(30 * time.Second)
		fmt.Printf("try to solve the riddle...\n")
		if ctf.Spec.Secretlocation == "Mount Myōboku" {
			fmt.Printf("you have reached Mount Myōboku ...\n")
			time.Sleep(30 * time.Second)
			fmt.Printf("You have found a scroll containing a secret ...\n")
			if err := r.createsecret(ctx); err != nil {
				return ctrl.Result{}, err
			}
			fmt.Printf("The secret has been revealed...\n")
			time.Sleep(30 * time.Second)
			if ctf.Spec.Secrectkey == "Ameterasu" {
				fmt.Printf("they found you ... \n")
				if err := r.createescapepod(ctx); err != nil {
					return ctrl.Result{}, err
				}
				time.Sleep(30 * time.Second)
				fmt.Printf("You need to escape now...\n")
			}
		} else {
			fmt.Printf("you didn't solve the riddle ...\n")
		}

	} else {
		fmt.Printf("Waiting for commands...\n")
	}

	return ctrl.Result{}, nil
}

func (r *CloudReconciler) ridlle() corev1.Pod {
	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "riddle",
			Namespace: "chunin",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "riddle",
					Image: "youffes/riddle",
				},
			},
		},
	}
	return pod
}
func (r *CloudReconciler) createriddle(ctx context.Context) error {
	pod := corev1.Pod{}
	podname := types.NamespacedName{
		Name:      "riddle",
		Namespace: "chunin",
	}
	if err := r.Get(ctx, podname, &pod); err != nil {
		if errors.IsNotFound(err) {
			pod := r.ridlle()
			if err := r.Create(ctx, &pod); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *CloudReconciler) secrect() corev1.Secret {
	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scroll",
			Namespace: "chunin",
		},
		StringData: map[string]string{
			"scroll": "hockage is weak now we need to attack now the keyword is Amaterasu",
		},
	}
	return secret
}
func (r *CloudReconciler) createsecret(ctx context.Context) error {
	secret := corev1.Secret{}
	secretname := types.NamespacedName{
		Name:      "scroll",
		Namespace: "chunin",
	}
	if err := r.Get(ctx, secretname, &secret); err != nil {
		if errors.IsNotFound(err) {
			secret := r.secrect()
			if err := r.Create(ctx, &secret); err != nil {
				return err
			}
		}
	}
	return nil
}
func (r *CloudReconciler) escapepod() corev1.Pod {
	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "escape",
			Namespace: "chunin",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "escape",
					Image: "youffes/escape",
				},
			},
		},
	}
	return pod
}

func (r *CloudReconciler) createescapepod(ctx context.Context) error {
	pod := corev1.Pod{}
	podname := types.NamespacedName{
		Name:      "escape",
		Namespace: "chunin",
	}
	if err := r.Get(ctx, podname, &pod); err != nil {
		if errors.IsNotFound(err) {
			pod := r.escapepod()
			if err := r.Create(ctx, &pod); err != nil {
				return err
			}
		}
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CloudReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ctfv1alpha1.Cloud{}).
		Complete(r)
}
