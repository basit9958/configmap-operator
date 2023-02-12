/*
Copyright 2023.

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

// DemoClusterConfigmapSpec defines the desired state of DemoClusterConfigmap
type DemoClusterConfigmapSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Data contains the configuration data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	Data map[string]string `json:"data"`
	// GenerateTo stores an array of LabelSelectors for matching on the specified namespace selectors.
	GenerateTo ClusterConfigMapGenerateTo `json:"generateTo"`
}

type ClusterConfigMapGenerateTo struct {
	NamespaceSelectors metav1.LabelSelector `json:"namespaceSelectors"`
}

// DemoClusterConfigmapStatus defines the observed state of DemoClusterConfigmap
type DemoClusterConfigmapStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ConfigMaps []corev1.ObjectReference `json:"configMaps,omitempty"`
	// ConfigMapsLock *sync.RWMutex            `json:"configMapsLock,omitempty"`
	// Namespaces []corev1.Namespace
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DemoClusterConfigmap is the Schema for the democlusterconfigmaps API
type DemoClusterConfigmap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoClusterConfigmapSpec   `json:"spec,omitempty"`
	Status DemoClusterConfigmapStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DemoClusterConfigmapList contains a list of DemoClusterConfigmap
type DemoClusterConfigmapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoClusterConfigmap `json:"items"`
	// ItemsLock       sync.RWMutex
}

func init() {
	SchemeBuilder.Register(&DemoClusterConfigmap{}, &DemoClusterConfigmapList{})
}
