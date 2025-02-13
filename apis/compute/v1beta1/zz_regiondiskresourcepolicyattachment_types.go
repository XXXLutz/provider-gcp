/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegionDiskResourcePolicyAttachmentObservation struct {

	// The name of the regional disk in which the resource policies are attached to.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// an identifier for the resource with format {{project}}/{{region}}/{{disk}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the region where the disk resides.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type RegionDiskResourcePolicyAttachmentParameters struct {

	// The name of the regional disk in which the resource policies are attached to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.RegionDisk
	// +kubebuilder:validation:Optional
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// Reference to a RegionDisk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskRef *v1.Reference `json:"diskRef,omitempty" tf:"-"`

	// Selector for a RegionDisk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskSelector *v1.Selector `json:"diskSelector,omitempty" tf:"-"`

	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ResourcePolicy
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a ResourcePolicy in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a ResourcePolicy in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the region where the disk resides.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// RegionDiskResourcePolicyAttachmentSpec defines the desired state of RegionDiskResourcePolicyAttachment
type RegionDiskResourcePolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionDiskResourcePolicyAttachmentParameters `json:"forProvider"`
}

// RegionDiskResourcePolicyAttachmentStatus defines the observed state of RegionDiskResourcePolicyAttachment.
type RegionDiskResourcePolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionDiskResourcePolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionDiskResourcePolicyAttachment is the Schema for the RegionDiskResourcePolicyAttachments API. Adds existing resource policies to a disk.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionDiskResourcePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   RegionDiskResourcePolicyAttachmentSpec   `json:"spec"`
	Status RegionDiskResourcePolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionDiskResourcePolicyAttachmentList contains a list of RegionDiskResourcePolicyAttachments
type RegionDiskResourcePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionDiskResourcePolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	RegionDiskResourcePolicyAttachment_Kind             = "RegionDiskResourcePolicyAttachment"
	RegionDiskResourcePolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionDiskResourcePolicyAttachment_Kind}.String()
	RegionDiskResourcePolicyAttachment_KindAPIVersion   = RegionDiskResourcePolicyAttachment_Kind + "." + CRDGroupVersion.String()
	RegionDiskResourcePolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(RegionDiskResourcePolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionDiskResourcePolicyAttachment{}, &RegionDiskResourcePolicyAttachmentList{})
}
