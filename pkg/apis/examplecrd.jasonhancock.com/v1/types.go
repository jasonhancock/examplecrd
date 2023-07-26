package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//go:generate controller-gen object paths=$GOFILE

// ExampleResourceSpec is the spec of the ExampleResource object.
// +kubebuilder:object:generate=true
type ExampleResourceSpec struct {
	Color string `json:"color"`
	Size  string `json:"size"`
}

// +genclient

// ExampleResource is our example resource.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ExampleResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ExampleResourceSpec `json:"spec"`
}

// ExampleResourceList is a list of ExampleResource objects.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ExampleResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ExampleResource `json:"items"`
}
