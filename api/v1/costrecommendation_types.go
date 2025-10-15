package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CostRecommendationSpec defines the desired state of CostRecommendation
type CostRecommendationSpec struct {
	ResourceType     string `json:"resourceType"`
	ResourceName     string `json:"resourceName"`
	Namespace        string `json:"namespace"`
	Issue            string `json:"issue"`
	Recommendation   string `json:"recommendation"`
	PotentialSavings string `json:"potentialSavings"`
	Confidence       string `json:"confidence"`
	AutoApplicable   bool   `json:"autoApplicable"`
}

// CostRecommendationStatus defines the observed state of CostRecommendation
type CostRecommendationStatus struct {
	Applied  bool   `json:"applied,omitempty"`
	Notes    string `json:"notes,omitempty"`
	LastSeen string `json:"lastSeen,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced

// CostRecommendation is a top-level CRD that holds a single recommendation
type CostRecommendation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CostRecommendationSpec   `json:"spec,omitempty"`
	Status CostRecommendationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CostRecommendationList contains a list of CostRecommendation
type CostRecommendationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CostRecommendation `json:"items"`
}

func init() {
	// NOTE: code-generation (controller-gen) should be run to generate DeepCopy methods
	// and register these types with a scheme. This placeholder keeps the types simple
	// so the scaffold can build without generated code.
}
