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

type InstanceIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type InstanceIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type InstanceIAMMemberObservation struct {
	Condition []InstanceIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type InstanceIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []InstanceIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +crossplane:generate:reference:type=Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// InstanceIAMMemberSpec defines the desired state of InstanceIAMMember
type InstanceIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceIAMMemberParameters `json:"forProvider"`
}

// InstanceIAMMemberStatus defines the observed state of InstanceIAMMember.
type InstanceIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIAMMember is the Schema for the InstanceIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.member)",message="member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.role)",message="role is a required parameter"
	Spec   InstanceIAMMemberSpec   `json:"spec"`
	Status InstanceIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIAMMemberList contains a list of InstanceIAMMembers
type InstanceIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceIAMMember `json:"items"`
}

// Repository type metadata.
var (
	InstanceIAMMember_Kind             = "InstanceIAMMember"
	InstanceIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceIAMMember_Kind}.String()
	InstanceIAMMember_KindAPIVersion   = InstanceIAMMember_Kind + "." + CRDGroupVersion.String()
	InstanceIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(InstanceIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceIAMMember{}, &InstanceIAMMemberList{})
}
