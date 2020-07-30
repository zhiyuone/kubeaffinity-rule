/*


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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AffinitySpec defines the desired state of Affinity
type AffinitySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Overwrite used to determine whether to overwrite the original rules
	Overwrite bool            `json:"overwrite,omitempty"`
	Affinity  corev1.Affinity `json:"affinity"`
}

// AffinityStatus defines the observed state of Affinity
type AffinityStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Affinity is the Schema for the affinities API
type Affinity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AffinitySpec   `json:"spec,omitempty"`
	Status AffinityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AffinityList contains a list of Affinity
type AffinityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Affinity `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Affinity{}, &AffinityList{})
}
