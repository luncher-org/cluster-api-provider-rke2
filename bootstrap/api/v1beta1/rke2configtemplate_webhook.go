/*
Copyright 2024 SUSE LLC.

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

package v1beta1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// RKE2configtemplatelog is for logging in this package.
var RKE2configtemplatelog = logf.Log.WithName("RKE2configtemplate-resource")

// SetupWebhookWithManager sets up and registers the webhook with the manager.
func (r *RKE2ConfigTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-bootstrap-cluster-x-k8s-io-v1beta1-rke2configtemplate,mutating=true,failurePolicy=fail,sideEffects=None,groups=bootstrap.cluster.x-k8s.io,resources=rke2configtemplates,verbs=create;update,versions=v1beta1,name=mrke2configtemplate.kb.io,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &RKE2ConfigTemplate{}

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (r *RKE2ConfigTemplate) Default(_ context.Context, _ runtime.Object) error {
	RKE2configtemplatelog.Info("default", "name", r.Name)

	return nil
}

//+kubebuilder:webhook:path=/validate-bootstrap-cluster-x-k8s-io-v1beta1-rke2configtemplate,mutating=false,failurePolicy=fail,sideEffects=None,groups=bootstrap.cluster.x-k8s.io,resources=rke2configtemplates,verbs=create;update,versions=v1beta1,name=vrke2configtemplate.kb.io,admissionReviewVersions=v1

var _ webhook.CustomValidator = &RKE2ConfigTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *RKE2ConfigTemplate) ValidateCreate(_ context.Context, _ runtime.Object) (admission.Warnings, error) {
	RKE2configtemplatelog.Info("validate create", "name", r.Name)

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *RKE2ConfigTemplate) ValidateUpdate(_ context.Context, _, _ runtime.Object) (admission.Warnings, error) {
	RKE2configtemplatelog.Info("validate update", "name", r.Name)

	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *RKE2ConfigTemplate) ValidateDelete(_ context.Context, _ runtime.Object) (admission.Warnings, error) {
	RKE2configtemplatelog.Info("validate delete", "name", r.Name)

	return nil, nil
}
