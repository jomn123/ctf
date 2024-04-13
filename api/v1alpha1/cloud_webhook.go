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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var cloudlog = logf.Log.WithName("cloud-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *Cloud) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-ctf-securinetes-com-v1alpha1-cloud,mutating=true,failurePolicy=fail,sideEffects=None,groups=ctf.securinetes.com,resources=clouds,verbs=create;update,versions=v1alpha1,name=mcloud.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Cloud{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Cloud) Default() {
	cloudlog.Info("default", "name", r.Name)
	if r.Spec.Intercept == "" {
		r.Spec.Intercept = ""
	}
	if r.Spec.Secretlocation == "" {
		r.Spec.Secretlocation = ""
	}
	if r.Spec.Secrectkey == "" {
		r.Spec.Secrectkey = ""
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-ctf-securinetes-com-v1alpha1-cloud,mutating=false,failurePolicy=fail,sideEffects=None,groups=ctf.securinetes.com,resources=clouds,verbs=create;update,versions=v1alpha1,name=vcloud.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Cloud{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Cloud) ValidateCreate() (admission.Warnings, error) {
	cloudlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Cloud) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	cloudlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Cloud) ValidateDelete() (admission.Warnings, error) {
	cloudlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
