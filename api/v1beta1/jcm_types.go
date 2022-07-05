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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// JcmSpec defines the desired state of Jcm
type JcmSpec struct {
	Prefix string `json:"prefix,omitempty"`
}

// JcmStatus defines the observed state of Jcm
type JcmStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase JcmPhase `json:"phase,omitempty"`

	Condition []JcmCondition `json:"conditions,omitempty"`

	Message string `json:"message,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
}

// JcmPhase is a label for the condition of a jcm at the current time.
type JcmPhase string

// These are the valid statuses of pods.
const (
	// JcmActivating is ready for use jcm
	JcmActivating JcmPhase = "Activated"
	// JcmTerminating is terminate jcm
	JcmTerminating JcmPhase = "Terminated"
	// JcmUnknown is before activating jcm
	JcmUnknown JcmPhase = "Unknown"
)

type JcmConditionType string

const (
	// Ready is initialize kv, db
	Ready JcmConditionType = "Ready"
)

// JcmCondition contains details for the current condition of this pod.
type JcmCondition struct {
	// Type is the type of the condition.
	Type JcmConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=PodConditionType"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Jcm is the Schema for the jcms API
type Jcm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JcmSpec   `json:"spec,omitempty"`
	Status JcmStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// JcmList contains a list of Jcm
type JcmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Jcm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Jcm{}, &JcmList{})
}
