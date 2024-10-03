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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TriggerParameters defines the desired state of Trigger
type TriggerParameters struct {
	// Region is which region the Trigger will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The actions initiated by this trigger when it fires.
	// +kubebuilder:validation:Required
	Actions []*Action `json:"actions"`
	// A description of the new trigger.
	Description *string `json:"description,omitempty"`
	// Batch condition that must be met (specified number of events received or
	// batch time window expired) before EventBridge event trigger fires.
	EventBatchingCondition *EventBatchingCondition `json:"eventBatchingCondition,omitempty"`
	// A predicate to specify when the new trigger should fire.
	//
	// This field is required when the trigger type is CONDITIONAL.
	Predicate *Predicate `json:"predicate,omitempty"`
	// A cron expression used to specify the schedule (see Time-Based Schedules
	// for Jobs and Crawlers (https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html).
	// For example, to run something every day at 12:15 UTC, you would specify:
	// cron(15 12 * * ? *).
	//
	// This field is required when the trigger type is SCHEDULED.
	Schedule *string `json:"schedule,omitempty"`
	// Set to true to start SCHEDULED and CONDITIONAL triggers when created. True
	// is not supported for ON_DEMAND triggers.
	StartOnCreation *bool `json:"startOnCreation,omitempty"`
	// The tags to use with this trigger. You may use tags to limit access to the
	// trigger. For more information about tags in Glue, see Amazon Web Services
	// Tags in Glue (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html)
	// in the developer guide.
	Tags map[string]*string `json:"tags,omitempty"`
	// The type of the new trigger.
	// +kubebuilder:validation:Required
	TriggerType *string `json:"triggerType"`
	// The name of the workflow associated with the trigger.
	WorkflowName            *string `json:"workflowName,omitempty"`
	CustomTriggerParameters `json:",inline"`
}

// TriggerSpec defines the desired state of Trigger
type TriggerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TriggerParameters `json:"forProvider"`
}

// TriggerObservation defines the observed state of Trigger
type TriggerObservation struct {
	// Reserved for future use.
	ID *string `json:"id,omitempty"`
	// The name of the trigger.
	Name *string `json:"name,omitempty"`
	// The current state of the trigger.
	State *string `json:"state,omitempty"`
}

// TriggerStatus defines the observed state of Trigger.
type TriggerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Trigger is the Schema for the Triggers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Trigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerSpec   `json:"spec"`
	Status            TriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerList contains a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trigger `json:"items"`
}

// Repository type metadata.
var (
	TriggerKind             = "Trigger"
	TriggerGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerKind}.String()
	TriggerKindAPIVersion   = TriggerKind + "." + GroupVersion.String()
	TriggerGroupVersionKind = GroupVersion.WithKind(TriggerKind)
)

func init() {
	SchemeBuilder.Register(&Trigger{}, &TriggerList{})
}
