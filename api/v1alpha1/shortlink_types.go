/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ShortLinkSpec defines the desired state of ShortLink
type ShortLinkSpec struct {
	// Alias is the short name (vanity name) of the shortening. If omitted, a random alias will be chosen
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=15
	Alias string `json:"alias"`

	// Target specifies the target to which we will redirect
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Target string `json:"target"`

	// RedirectAfter specifies after how many seconds to redirect (Default=3)
	// +kubebuilder:default:=3
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=99
	RedirectAfter int64 `json:"after,omitempty"`
}

// ShortLinkStatus defines the observed state of ShortLink
type ShortLinkStatus struct {
	// Count represents how often this ShortLink has been called
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	Count int `json:"count"`

	// Ready indicates if the shortlink is ready to be consumed (all labels, etc. are set)
	// +kubebuilder:default:=false
	Ready bool `json:"ready,omitempty"`
}

// ShortLink is the Schema for the shortlinks API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Alias",type=string,JSONPath=`.spec.alias`
// +kubebuilder:printcolumn:name="Target",type=string,JSONPath=`.spec.target`
// +kubebuilder:printcolumn:name="After",type=string,JSONPath=`.spec.after`
// +kubebuilder:printcolumn:name="Invoked",type=string,JSONPath=`.status.count`
// +k8s:openapi-gen=true
type ShortLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShortLinkSpec   `json:"spec,omitempty"`
	Status ShortLinkStatus `json:"status,omitempty"`
}

// ShortLinkList contains a list of ShortLink
// +kubebuilder:object:root=true
type ShortLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShortLink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ShortLink{}, &ShortLinkList{})
}
