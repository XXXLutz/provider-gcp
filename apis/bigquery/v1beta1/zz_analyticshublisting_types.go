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

type AnalyticsHubListingObservation struct {

	// Shared dataset i.e. BigQuery dataset source.
	// Structure is documented below.
	BigqueryDataset []BigqueryDatasetObservation `json:"bigqueryDataset,omitempty" tf:"bigquery_dataset,omitempty"`

	// Categories of the listing. Up to two categories are allowed.
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeID *string `json:"dataExchangeId,omitempty" tf:"data_exchange_id,omitempty"`

	// Details of the data provider who owns the source data.
	// Structure is documented below.
	DataProvider []DataProviderObservation `json:"dataProvider,omitempty" tf:"data_provider,omitempty"`

	// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Documentation describing the listing.
	Documentation *string `json:"documentation,omitempty" tf:"documentation,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Base64 encoded image representing the listing.
	Icon *string `json:"icon,omitempty" tf:"icon,omitempty"`

	// The name of the location this data exchange listing.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name of the listing. e.g. "projects/myproject/locations/US/dataExchanges/123/listings/456"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the listing publisher.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Details of the publisher who owns the listing and who can share the source data.
	// Structure is documented below.
	Publisher []PublisherObservation `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
	RequestAccess *string `json:"requestAccess,omitempty" tf:"request_access,omitempty"`
}

type AnalyticsHubListingParameters struct {

	// Shared dataset i.e. BigQuery dataset source.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BigqueryDataset []BigqueryDatasetParameters `json:"bigqueryDataset,omitempty" tf:"bigquery_dataset,omitempty"`

	// Categories of the listing. Up to two categories are allowed.
	// +kubebuilder:validation:Optional
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.AnalyticsHubDataExchange
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("data_exchange_id",false)
	// +kubebuilder:validation:Optional
	DataExchangeID *string `json:"dataExchangeId,omitempty" tf:"data_exchange_id,omitempty"`

	// Reference to a AnalyticsHubDataExchange in bigquery to populate dataExchangeId.
	// +kubebuilder:validation:Optional
	DataExchangeIDRef *v1.Reference `json:"dataExchangeIdRef,omitempty" tf:"-"`

	// Selector for a AnalyticsHubDataExchange in bigquery to populate dataExchangeId.
	// +kubebuilder:validation:Optional
	DataExchangeIDSelector *v1.Selector `json:"dataExchangeIdSelector,omitempty" tf:"-"`

	// Details of the data provider who owns the source data.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DataProvider []DataProviderParameters `json:"dataProvider,omitempty" tf:"data_provider,omitempty"`

	// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Documentation describing the listing.
	// +kubebuilder:validation:Optional
	Documentation *string `json:"documentation,omitempty" tf:"documentation,omitempty"`

	// Base64 encoded image representing the listing.
	// +kubebuilder:validation:Optional
	Icon *string `json:"icon,omitempty" tf:"icon,omitempty"`

	// The name of the location this data exchange listing.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Email or URL of the listing publisher.
	// +kubebuilder:validation:Optional
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Details of the publisher who owns the listing and who can share the source data.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Publisher []PublisherParameters `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
	// +kubebuilder:validation:Optional
	RequestAccess *string `json:"requestAccess,omitempty" tf:"request_access,omitempty"`
}

type BigqueryDatasetObservation struct {

	// Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`
}

type BigqueryDatasetParameters struct {

	// Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// Reference to a Dataset in bigquery to populate dataset.
	// +kubebuilder:validation:Optional
	DatasetRef *v1.Reference `json:"datasetRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate dataset.
	// +kubebuilder:validation:Optional
	DatasetSelector *v1.Selector `json:"datasetSelector,omitempty" tf:"-"`
}

type DataProviderObservation struct {

	// Name of the data provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the data provider.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type DataProviderParameters struct {

	// Name of the data provider.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Email or URL of the data provider.
	// +kubebuilder:validation:Optional
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type PublisherObservation struct {

	// Name of the listing publisher.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the listing publisher.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type PublisherParameters struct {

	// Name of the listing publisher.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Email or URL of the listing publisher.
	// +kubebuilder:validation:Optional
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

// AnalyticsHubListingSpec defines the desired state of AnalyticsHubListing
type AnalyticsHubListingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyticsHubListingParameters `json:"forProvider"`
}

// AnalyticsHubListingStatus defines the observed state of AnalyticsHubListing.
type AnalyticsHubListingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyticsHubListingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsHubListing is the Schema for the AnalyticsHubListings API. A Bigquery Analytics Hub data exchange listing
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AnalyticsHubListing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.bigqueryDataset)",message="bigqueryDataset is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   AnalyticsHubListingSpec   `json:"spec"`
	Status AnalyticsHubListingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsHubListingList contains a list of AnalyticsHubListings
type AnalyticsHubListingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsHubListing `json:"items"`
}

// Repository type metadata.
var (
	AnalyticsHubListing_Kind             = "AnalyticsHubListing"
	AnalyticsHubListing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnalyticsHubListing_Kind}.String()
	AnalyticsHubListing_KindAPIVersion   = AnalyticsHubListing_Kind + "." + CRDGroupVersion.String()
	AnalyticsHubListing_GroupVersionKind = CRDGroupVersion.WithKind(AnalyticsHubListing_Kind)
)

func init() {
	SchemeBuilder.Register(&AnalyticsHubListing{}, &AnalyticsHubListingList{})
}
