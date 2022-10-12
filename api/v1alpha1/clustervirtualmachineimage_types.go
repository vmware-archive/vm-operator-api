// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (clusterVirtualMachineImage *ClusterVirtualMachineImage) GetConditions() Conditions {
	return clusterVirtualMachineImage.Status.Conditions
}

func (clusterVirtualMachineImage *ClusterVirtualMachineImage) SetConditions(conditions Conditions) {
	clusterVirtualMachineImage.Status.Conditions = conditions
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=cvmi;cvmimage;clustervmimage
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ContentLibraryName",type="string",JSONPath=".status.contentLibraryRef.name"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.productInfo.version"
// +kubebuilder:printcolumn:name="OsType",type="string",JSONPath=".spec.osInfo.type"
// +kubebuilder:printcolumn:name="Format",type="string",JSONPath=".spec.type"
// +kubebuilder:printcolumn:name="ImageSupported",type="boolean",priority=1,JSONPath=".status.imageSupported"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ClusterVirtualMachineImage is the schema for the clustervirtualmachineimage API
// A ClusterVirtualMachineImage represents the desired specification and the observed status of a
// ClusterVirtualMachineImage instance.
type ClusterVirtualMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineImageSpec   `json:"spec,omitempty"`
	Status VirtualMachineImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterVirtualMachineImageList contains a list of ClusterVirtualMachineImage.
type ClusterVirtualMachineImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterVirtualMachineImage `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&ClusterVirtualMachineImage{}, &ClusterVirtualMachineImageList{})
}
