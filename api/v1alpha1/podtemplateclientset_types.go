package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodTemplateClientSetSpec defines the desired state of PodTemplateClientSet
type PodTemplateClientSetSpec struct {
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	Template corev1.PodTemplateSpec `json:"template"`
}

// PodTemplateClientSetStatus defines the observed state of PodTemplateClientSet
type PodTemplateClientSetStatus struct {
	// Current replicas, integrated in /scale subresource
	Replicas int32 `json:"replicas"`

	// Selector for pods belonging to this ClientSet, integrated in /scale subresource
	Selector string `json:"selector"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=ptcs
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.status.selector
// +kubebuilder:printcolumn:JSONPath=".status.replicas",name=Replicas,type=string

// PodTemplateClientSet is the Schema for the podtemplateclientsets API
type PodTemplateClientSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodTemplateClientSetSpec   `json:"spec,omitempty"`
	Status PodTemplateClientSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodTemplateClientSetList contains a list of PodTemplateClientSet
type PodTemplateClientSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodTemplateClientSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodTemplateClientSet{}, &PodTemplateClientSetList{})
}
