// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterVirtualMachineImageSpec defines the desired state of ClusterVirtualMachineImageSpec
type ClusterVirtualMachineImageSpec struct {
	// Type describes the type of the ClusterVirtualMachineImage. Currently, the only supported image is "OVF"
	Type string `json:"type"`

	// ImageSourceType describes the type of content source of the ClusterVirtualMachineImage. The only Content Source
	// supported currently is the vSphere Content Library
	// +optional
	ImageSourceType string `json:"imageSourceType,omitempty"`

	// ImageID is a unique identifier exposed by the provider of this ClusterVirtualMachineImage.
	ImageID string `json:"imageID"`

	// TODO: ProviderRef

	// ProductInfo describes the attributes of the VirtualMachineImage relating to the product contained in the
	// image.
	// +optional
	ProductInfo VirtualMachineImageProductInfo `json:"productInfo,omitempty"`

	// OSInfo describes the attributes of the VirtualMachineImage relating to the Operating System contained in the
	// image.
	// +optional
	OSInfo VirtualMachineImageOSInfo `json:"osInfo,omitempty"`

	// OVFEnv describes the user configurable customization parameters of the ClusterVirtualMachineImage.
	// +optional
	OVFEnv map[string]OvfProperty `json:"ovfEnv,omitempty"`

	// HardwareVersion describes the virtual hardware version of the image
	// +optional
	HardwareVersion int32 `json:"hwVersion,omitempty"`
}

// ClusterVirtualMachineImageStatus defines the observed state of ClusterVirtualMachineImage
type ClusterVirtualMachineImageStatus struct {
	// TODO: Deprecated fields

	// ImageName describes the display name of this ClusterVirtualMachineImage.
	// +optional
	ImageName string `json:"imageName,omitempty"`

	// ImageSupported indicates whether the ClusterVirtualMachineImage is supported by VMService.
	// A ClusterVirtualMachineImage is supported by VMService if the following conditions are true:
	// - VirtualMachineImageV1Alpha1CompatibleCondition
	// +optional
	ImageSupported *bool `json:"imageSupported,omitempty"`

	// Conditions describes the current condition information of the ClusterVirtualMachineImage object.
	// e.g. if the OS type is supported or image is supported by VMService
	// +optional
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

func (clusterVirtualMachineImage *ClusterVirtualMachineImage) GetConditions() Conditions {
	return clusterVirtualMachineImage.Status.Conditions
}

func (clusterVirtualMachineImage *ClusterVirtualMachineImage) SetConditions(conditions Conditions) {
	clusterVirtualMachineImage.Status.Conditions = conditions
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=vmimage
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// Todo:
// +kubebuilder:printcolumn:name="ContentSourceName",type="string",JSONPath=".spec.providerRef.name"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.productInfo.version"
// +kubebuilder:printcolumn:name="OsType",type="string",JSONPath=".spec.osInfo.type"
// +kubebuilder:printcolumn:name="Format",type="string",JSONPath=".spec.type"
// +kubebuilder:printcolumn:name="ImageSupported",type="boolean",priority=1,JSONPath=".status.imageSupported"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ClusterVirtualMachineImage is the schema for the clustervirtualmachineimage API
// A ClusterVirtualMachineImage represents the desired specification and the observed status of a
// ClusterVirtualMachineImage instance
type ClusterVirtualMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterVirtualMachineImageSpec   `json:"spec,omitempty"`
	Status ClusterVirtualMachineImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterVirtualMachineImageList contains a list of ClusterVirtualMachineImage
type ClusterVirtualMachineImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterVirtualMachineImage `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&ClusterVirtualMachineImage{}, &ClusterVirtualMachineImageList{})
}
