/*
Copyright IBM Corp. 2023
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitorSpec defines the desired state of Monitor
type MonitorSpec struct {
	// Image is the Image to use
	Image string `json:"image,omitempty"`

	// Tag is the image tag to use
	Tag string `json:"tag,omitempty"`

	// Name/Value array used when starting the Job
	Name []string `json:"name_value,omitempty"`

	// Path in the container to the Ansible YAML that is to be executed
	Path string `json:"path,omitempty"`

	// Recovery Image is the Image to use
	RecoveryImage string `json:"recovery_image,omitempty"`

	// Recovery Tag is the image tag to use
	RecoveryTag string `json:"recovery_tag,omitempty"`

	// Recovery Name/Value array used when starting the Job
	RecoveryName []string `json:"recovery_name_value,omitempty"`

	// Recovery Path in the container to the Ansible YAML that is to be executed
	RecoveryPath string `json:"recovery_path,omitempty"`

	// The Deployment Name
	Deployment string `json:"deployment_name,omitempty"`
}

// MonitorStatus defines the observed state of Monitor
type MonitorStatus struct {
	// The status of the running deployment
	Status string `json:"status,omitempty"`

	// The status of the job recovering the deployment
	RecoveryStatus string `json:"recovery_status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Monitor is the Schema for the monitors API
type Monitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitorSpec   `json:"spec,omitempty"`
	Status MonitorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MonitorList contains a list of Monitor
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Monitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Monitor{}, &MonitorList{})
}
