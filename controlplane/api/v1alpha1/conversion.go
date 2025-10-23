/*
Copyright 2023 SUSE.

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
	"fmt"

	apiconversion "k8s.io/apimachinery/pkg/conversion"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	bootstrapv1alpha1 "github.com/rancher/cluster-api-provider-rke2/bootstrap/api/v1alpha1"
	bootstrapv1beta1 "github.com/rancher/cluster-api-provider-rke2/bootstrap/api/v1beta1"
	controlplanev1 "github.com/rancher/cluster-api-provider-rke2/controlplane/api/v1beta1"
)

func (src *RKE2ControlPlane) ConvertTo(dstRaw conversion.Hub) error {
	dst, ok := dstRaw.(*controlplanev1.RKE2ControlPlane)
	if !ok {
		return fmt.Errorf("not a RKE2ControlPlane: %v", dst)
	}
	if err := Convert_v1alpha1_RKE2ControlPlane_To_v1beta1_RKE2ControlPlane(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.Version = src.Spec.AgentConfig.Version

	// Manually restore data.
	restored := &controlplanev1.RKE2ControlPlane{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	if restored.Spec.GzipUserData != nil {
		dst.Spec.GzipUserData = restored.Spec.GzipUserData
	}

	if restored.Spec.Version != "" {
		dst.Spec.Version = restored.Spec.Version
	}

	if restored.Spec.AgentConfig.AirGappedChecksum != "" {
		dst.Spec.AgentConfig.AirGappedChecksum = restored.Spec.AgentConfig.AirGappedChecksum
	}

	if restored.Spec.AgentConfig.PodSecurityAdmissionConfigFile != "" {
		dst.Spec.AgentConfig.PodSecurityAdmissionConfigFile = restored.Spec.AgentConfig.PodSecurityAdmissionConfigFile
	}

	if restored.Spec.RemediationStrategy != nil {
		dst.Spec.RemediationStrategy = restored.Spec.RemediationStrategy
	}

	if restored.Spec.ServerConfig.ExternalDatastoreSecret != nil {
		dst.Spec.ServerConfig.ExternalDatastoreSecret = restored.Spec.ServerConfig.ExternalDatastoreSecret
	}

	dst.Spec.ServerConfig.EmbeddedRegistry = restored.Spec.ServerConfig.EmbeddedRegistry
	dst.Spec.ServerConfig.SecretsEncryptionProvider = restored.Spec.ServerConfig.SecretsEncryptionProvider
	dst.Spec.MachineTemplate = restored.Spec.MachineTemplate
	dst.Status = restored.Status
	dst.Spec.Files = restored.Spec.Files

	return nil
}

func (dst *RKE2ControlPlane) ConvertFrom(srcRaw conversion.Hub) error {
	src, ok := srcRaw.(*controlplanev1.RKE2ControlPlane)
	if !ok {
		return fmt.Errorf("not a RKE2ControlPlane: %v", src)
	}

	if err := Convert_v1beta1_RKE2ControlPlane_To_v1alpha1_RKE2ControlPlane(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.AgentConfig.Version = src.Spec.Version

	// Preserve `NodeDeletionTimeout` since it does not exist in v1alpha1
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

func (src *RKE2ControlPlaneList) ConvertTo(dstRaw conversion.Hub) error {
	dst, ok := dstRaw.(*controlplanev1.RKE2ControlPlaneList)
	if !ok {
		return fmt.Errorf("not a RKE2ControlPlaneList: %v", dst)
	}

	if err := Convert_v1alpha1_RKE2ControlPlaneList_To_v1beta1_RKE2ControlPlaneList(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (dst *RKE2ControlPlaneList) ConvertFrom(srcRaw conversion.Hub) error {
	src, ok := srcRaw.(*controlplanev1.RKE2ControlPlaneList)
	if !ok {
		return fmt.Errorf("not a RKE2ControlPlaneList: %v", src)
	}

	if err := Convert_v1beta1_RKE2ControlPlaneList_To_v1alpha1_RKE2ControlPlaneList(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *RKE2ControlPlaneTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst, ok := dstRaw.(*controlplanev1.RKE2ControlPlaneTemplate)
	if !ok {
		return fmt.Errorf("not a RKE2ControlPlaneTemplate: %v", dst)
	}

	if err := Convert_v1alpha1_RKE2ControlPlaneTemplate_To_v1beta1_RKE2ControlPlaneTemplate(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &controlplanev1.RKE2ControlPlaneTemplate{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	if restored.Spec.Template.Spec.AgentConfig.AirGappedChecksum != "" {
		dst.Spec.Template.Spec.AgentConfig.AirGappedChecksum = restored.Spec.Template.Spec.AgentConfig.AirGappedChecksum
	}

	if restored.Spec.Template.Spec.AgentConfig.PodSecurityAdmissionConfigFile != "" {
		dst.Spec.Template.Spec.AgentConfig.PodSecurityAdmissionConfigFile = restored.Spec.Template.Spec.AgentConfig.PodSecurityAdmissionConfigFile
	}

	dst.Spec.Template.Spec.ServerConfig.EmbeddedRegistry = restored.Spec.Template.Spec.ServerConfig.EmbeddedRegistry
	dst.Spec.Template.Spec.ServerConfig.SecretsEncryptionProvider = restored.Spec.Template.Spec.ServerConfig.SecretsEncryptionProvider
	dst.Spec.Template = restored.Spec.Template
	dst.Status = restored.Status
	dst.Spec.Template.Spec.MachineTemplate.NodeDrainTimeout = restored.Spec.Template.Spec.MachineTemplate.NodeDrainTimeout
	dst.Spec.Template.Spec.MachineTemplate.NodeDeletionTimeout = restored.Spec.Template.Spec.MachineTemplate.NodeDeletionTimeout
	dst.Spec.Template.Spec.MachineTemplate.NodeVolumeDetachTimeout = restored.Spec.Template.Spec.MachineTemplate.NodeVolumeDetachTimeout

	dst.Spec.Template.Spec.Files = restored.Spec.Template.Spec.Files
	return nil
}

func (dst *RKE2ControlPlaneTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src, ok := srcRaw.(*controlplanev1.RKE2ControlPlaneTemplate)
	if !ok {
		return fmt.Errorf("not a RKE2ControlPlaneTemplate: %v", src)
	}

	if err := Convert_v1beta1_RKE2ControlPlaneTemplate_To_v1alpha1_RKE2ControlPlaneTemplate(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

func (src *RKE2ControlPlaneTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst, ok := dstRaw.(*controlplanev1.RKE2ControlPlaneTemplateList)
	if !ok {
		return fmt.Errorf("not a RKE2ControlPlaneTemplateList: %v", dst)
	}

	if err := Convert_v1alpha1_RKE2ControlPlaneTemplateList_To_v1beta1_RKE2ControlPlaneTemplateList(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (dst *RKE2ControlPlaneTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src, ok := srcRaw.(*controlplanev1.RKE2ControlPlaneTemplateList)
	if !ok {
		return fmt.Errorf("not a RKE2ControlPlaneTemplateList: %v", src)
	}

	if err := Convert_v1beta1_RKE2ControlPlaneTemplateList_To_v1alpha1_RKE2ControlPlaneTemplateList(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func Convert_v1beta1_RKE2ControlPlaneSpec_To_v1alpha1_RKE2ControlPlaneSpec(in *controlplanev1.RKE2ControlPlaneSpec, out *RKE2ControlPlaneSpec, s apiconversion.Scope) error {
	// Version was added in v1beta1.
	// MachineTemplate was added in v1beta1.
	return autoConvert_v1beta1_RKE2ControlPlaneSpec_To_v1alpha1_RKE2ControlPlaneSpec(in, out, s)
}

func Convert_v1beta1_RKE2ControlPlaneStatus_To_v1alpha1_RKE2ControlPlaneStatus(in *controlplanev1.RKE2ControlPlaneStatus, out *RKE2ControlPlaneStatus, s apiconversion.Scope) error {
	return autoConvert_v1beta1_RKE2ControlPlaneStatus_To_v1alpha1_RKE2ControlPlaneStatus(in, out, s)
}

func Convert_v1alpha1_RKE2ControlPlaneStatus_To_v1beta1_RKE2ControlPlaneStatus(in *RKE2ControlPlaneStatus, out *controlplanev1.RKE2ControlPlaneStatus, s apiconversion.Scope) error {
	return autoConvert_v1alpha1_RKE2ControlPlaneStatus_To_v1beta1_RKE2ControlPlaneStatus(in, out, s)
}

func Convert_v1beta1_RKE2ControlPlaneTemplateSpec_To_v1alpha1_RKE2ControlPlaneTemplateSpec(in *controlplanev1.RKE2ControlPlaneTemplateSpec, out *RKE2ControlPlaneTemplateSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_RKE2ControlPlaneTemplateSpec_To_v1alpha1_RKE2ControlPlaneTemplateSpec(in, out, s)
}

func Convert_v1alpha1_RKE2ControlPlaneTemplateSpec_To_v1beta1_RKE2ControlPlaneTemplateSpec(in *RKE2ControlPlaneTemplateSpec, out *controlplanev1.RKE2ControlPlaneTemplateSpec, s apiconversion.Scope) error {
	return autoConvert_v1alpha1_RKE2ControlPlaneTemplateSpec_To_v1beta1_RKE2ControlPlaneTemplateSpec(in, out, s)
}

func Convert_v1alpha1_RKE2ControlPlaneTemplateStatus_To_v1beta1_RKE2ControlPlaneStatus(in *RKE2ControlPlaneTemplateStatus, out *controlplanev1.RKE2ControlPlaneStatus, s apiconversion.Scope) error {
	return nil
}

func Convert_v1beta1_RKE2ControlPlaneStatus_To_v1alpha1_RKE2ControlPlaneTemplateStatus(in *controlplanev1.RKE2ControlPlaneStatus, out *RKE2ControlPlaneTemplateStatus, s apiconversion.Scope) error {
	return nil
}

func Convert_v1beta1_RKE2ConfigSpec_To_v1alpha1_RKE2ConfigSpec(in *bootstrapv1beta1.RKE2ConfigSpec, out *bootstrapv1alpha1.RKE2ConfigSpec, s apiconversion.Scope) error {
	return bootstrapv1alpha1.Convert_v1beta1_RKE2ConfigSpec_To_v1alpha1_RKE2ConfigSpec(in, out, s)
}

func Convert_v1alpha1_RKE2ConfigSpec_To_v1beta1_RKE2ConfigSpec(in *bootstrapv1alpha1.RKE2ConfigSpec, out *bootstrapv1beta1.RKE2ConfigSpec, s apiconversion.Scope) error {
	return bootstrapv1alpha1.Convert_v1alpha1_RKE2ConfigSpec_To_v1beta1_RKE2ConfigSpec(in, out, s)
}

func Convert_v1beta1_RKE2ServerConfig_To_v1alpha1_RKE2ServerConfig(in *controlplanev1.RKE2ServerConfig, out *RKE2ServerConfig, s apiconversion.Scope) error {
	return autoConvert_v1beta1_RKE2ServerConfig_To_v1alpha1_RKE2ServerConfig(in, out, s)
}
