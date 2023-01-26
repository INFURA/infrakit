package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodTemplateClientSetSpec defines the desired state of PodTemplateClientSet
type PodTemplateClientSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PodTemplateClientSet. Edit podtemplateclientset_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// PodTemplateClientSetStatus defines the observed state of PodTemplateClientSet
type PodTemplateClientSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

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
