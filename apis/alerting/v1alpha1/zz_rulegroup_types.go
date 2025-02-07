// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DataInitParameters struct {

	// 100" if this stage is an expression stage.
	// The UID of the datasource being queried, or "-100" if this stage is an expression stage.
	DatasourceUID *string `json:"datasourceUid,omitempty" tf:"datasource_uid,omitempty"`

	// (String) Custom JSON data to send to the specified datasource when querying.
	// Custom JSON data to send to the specified datasource when querying.
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// (String) An optional identifier for the type of query being executed. Defaults to “.
	// An optional identifier for the type of query being executed. Defaults to “.
	QueryType *string `json:"queryType,omitempty" tf:"query_type,omitempty"`

	// (String) A unique string to identify this query stage within a rule.
	// A unique string to identify this query stage within a rule.
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`

	// (Block List, Min: 1, Max: 1) The time range, relative to when the query is executed, across which to query. (see below for nested schema)
	// The time range, relative to when the query is executed, across which to query.
	RelativeTimeRange []RelativeTimeRangeInitParameters `json:"relativeTimeRange,omitempty" tf:"relative_time_range,omitempty"`
}

type DataObservation struct {

	// 100" if this stage is an expression stage.
	// The UID of the datasource being queried, or "-100" if this stage is an expression stage.
	DatasourceUID *string `json:"datasourceUid,omitempty" tf:"datasource_uid,omitempty"`

	// (String) Custom JSON data to send to the specified datasource when querying.
	// Custom JSON data to send to the specified datasource when querying.
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// (String) An optional identifier for the type of query being executed. Defaults to “.
	// An optional identifier for the type of query being executed. Defaults to “.
	QueryType *string `json:"queryType,omitempty" tf:"query_type,omitempty"`

	// (String) A unique string to identify this query stage within a rule.
	// A unique string to identify this query stage within a rule.
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`

	// (Block List, Min: 1, Max: 1) The time range, relative to when the query is executed, across which to query. (see below for nested schema)
	// The time range, relative to when the query is executed, across which to query.
	RelativeTimeRange []RelativeTimeRangeObservation `json:"relativeTimeRange,omitempty" tf:"relative_time_range,omitempty"`
}

type DataParameters struct {

	// 100" if this stage is an expression stage.
	// The UID of the datasource being queried, or "-100" if this stage is an expression stage.
	// +kubebuilder:validation:Optional
	DatasourceUID *string `json:"datasourceUid" tf:"datasource_uid,omitempty"`

	// (String) Custom JSON data to send to the specified datasource when querying.
	// Custom JSON data to send to the specified datasource when querying.
	// +kubebuilder:validation:Optional
	Model *string `json:"model" tf:"model,omitempty"`

	// (String) An optional identifier for the type of query being executed. Defaults to “.
	// An optional identifier for the type of query being executed. Defaults to “.
	// +kubebuilder:validation:Optional
	QueryType *string `json:"queryType,omitempty" tf:"query_type,omitempty"`

	// (String) A unique string to identify this query stage within a rule.
	// A unique string to identify this query stage within a rule.
	// +kubebuilder:validation:Optional
	RefID *string `json:"refId" tf:"ref_id,omitempty"`

	// (Block List, Min: 1, Max: 1) The time range, relative to when the query is executed, across which to query. (see below for nested schema)
	// The time range, relative to when the query is executed, across which to query.
	// +kubebuilder:validation:Optional
	RelativeTimeRange []RelativeTimeRangeParameters `json:"relativeTimeRange" tf:"relative_time_range,omitempty"`
}

type RelativeTimeRangeInitParameters struct {

	// (Number) The number of seconds in the past, relative to when the rule is evaluated, at which the time range begins.
	// The number of seconds in the past, relative to when the rule is evaluated, at which the time range begins.
	From *float64 `json:"from,omitempty" tf:"from,omitempty"`

	// (Number) The number of seconds in the past, relative to when the rule is evaluated, at which the time range ends.
	// The number of seconds in the past, relative to when the rule is evaluated, at which the time range ends.
	To *float64 `json:"to,omitempty" tf:"to,omitempty"`
}

type RelativeTimeRangeObservation struct {

	// (Number) The number of seconds in the past, relative to when the rule is evaluated, at which the time range begins.
	// The number of seconds in the past, relative to when the rule is evaluated, at which the time range begins.
	From *float64 `json:"from,omitempty" tf:"from,omitempty"`

	// (Number) The number of seconds in the past, relative to when the rule is evaluated, at which the time range ends.
	// The number of seconds in the past, relative to when the rule is evaluated, at which the time range ends.
	To *float64 `json:"to,omitempty" tf:"to,omitempty"`
}

type RelativeTimeRangeParameters struct {

	// (Number) The number of seconds in the past, relative to when the rule is evaluated, at which the time range begins.
	// The number of seconds in the past, relative to when the rule is evaluated, at which the time range begins.
	// +kubebuilder:validation:Optional
	From *float64 `json:"from" tf:"from,omitempty"`

	// (Number) The number of seconds in the past, relative to when the rule is evaluated, at which the time range ends.
	// The number of seconds in the past, relative to when the rule is evaluated, at which the time range ends.
	// +kubebuilder:validation:Optional
	To *float64 `json:"to" tf:"to,omitempty"`
}

type RuleGroupInitParameters struct {

	// Defaults to false. Defaults to `false`.
	DisableProvenance *bool `json:"disableProvenance,omitempty" tf:"disable_provenance,omitempty"`

	// (Number) The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
	// The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
	IntervalSeconds *float64 `json:"intervalSeconds,omitempty" tf:"interval_seconds,omitempty"`

	// (String) The name of the rule group.
	// The name of the rule group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1) The rules within the group. (see below for nested schema)
	// The rules within the group.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RuleGroupObservation struct {

	// Defaults to false. Defaults to `false`.
	DisableProvenance *bool `json:"disableProvenance,omitempty" tf:"disable_provenance,omitempty"`

	// (String) The UID of the folder that the group belongs to.
	// The UID of the folder that the group belongs to.
	FolderUID *string `json:"folderUid,omitempty" tf:"folder_uid,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
	// The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
	IntervalSeconds *float64 `json:"intervalSeconds,omitempty" tf:"interval_seconds,omitempty"`

	// (String) The name of the rule group.
	// The name of the rule group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (Block List, Min: 1) The rules within the group. (see below for nested schema)
	// The rules within the group.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RuleGroupParameters struct {

	// Defaults to false. Defaults to `false`.
	// +kubebuilder:validation:Optional
	DisableProvenance *bool `json:"disableProvenance,omitempty" tf:"disable_provenance,omitempty"`

	// Reference to a Folder in oss to populate folderUid.
	// +kubebuilder:validation:Optional
	FolderRef *v1.Reference `json:"folderRef,omitempty" tf:"-"`

	// Selector for a Folder in oss to populate folderUid.
	// +kubebuilder:validation:Optional
	FolderSelector *v1.Selector `json:"folderSelector,omitempty" tf:"-"`

	// (String) The UID of the folder that the group belongs to.
	// The UID of the folder that the group belongs to.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Folder
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.UIDExtractor()
	// +crossplane:generate:reference:refFieldName=FolderRef
	// +crossplane:generate:reference:selectorFieldName=FolderSelector
	// +kubebuilder:validation:Optional
	FolderUID *string `json:"folderUid,omitempty" tf:"folder_uid,omitempty"`

	// (Number) The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
	// The interval, in seconds, at which all rules in the group are evaluated. If a group contains many rules, the rules are evaluated sequentially.
	// +kubebuilder:validation:Optional
	IntervalSeconds *float64 `json:"intervalSeconds,omitempty" tf:"interval_seconds,omitempty"`

	// (String) The name of the rule group.
	// The name of the rule group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +crossplane:generate:reference:selectorFieldName=OrganizationSelector
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`

	// (Block List, Min: 1) The rules within the group. (see below for nested schema)
	// The rules within the group.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RuleInitParameters struct {

	// value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to map[].
	// Key-value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to `map[]`.
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// (String) The ref_id of the query node in the data field to use as the alert condition.
	// The `ref_id` of the query node in the `data` field to use as the alert condition.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// (Block List, Min: 1) A sequence of stages that describe the contents of the rule. (see below for nested schema)
	// A sequence of stages that describe the contents of the rule.
	Data []DataInitParameters `json:"data,omitempty" tf:"data,omitempty"`

	// (String) Describes what state to enter when the rule's query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to Alerting.
	// Describes what state to enter when the rule's query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to `Alerting`.
	ExecErrState *string `json:"execErrState,omitempty" tf:"exec_err_state,omitempty"`

	// (String) The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to 0.
	// The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
	For *string `json:"for,omitempty" tf:"for,omitempty"`

	// (Boolean) Sets whether the alert should be paused or not. Defaults to false.
	// Sets whether the alert should be paused or not. Defaults to `false`.
	IsPaused *bool `json:"isPaused,omitempty" tf:"is_paused,omitempty"`

	// value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to map[].
	// Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the rule group.
	// The name of the alert rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Describes what state to enter when the rule's query returns No Data. Options are OK, NoData, and Alerting. Defaults to NoData.
	// Describes what state to enter when the rule's query returns No Data. Options are OK, NoData, and Alerting. Defaults to `NoData`.
	NoDataState *string `json:"noDataState,omitempty" tf:"no_data_state,omitempty"`
}

type RuleObservation struct {

	// value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to map[].
	// Key-value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to `map[]`.
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// (String) The ref_id of the query node in the data field to use as the alert condition.
	// The `ref_id` of the query node in the `data` field to use as the alert condition.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// (Block List, Min: 1) A sequence of stages that describe the contents of the rule. (see below for nested schema)
	// A sequence of stages that describe the contents of the rule.
	Data []DataObservation `json:"data,omitempty" tf:"data,omitempty"`

	// (String) Describes what state to enter when the rule's query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to Alerting.
	// Describes what state to enter when the rule's query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to `Alerting`.
	ExecErrState *string `json:"execErrState,omitempty" tf:"exec_err_state,omitempty"`

	// (String) The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to 0.
	// The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
	For *string `json:"for,omitempty" tf:"for,omitempty"`

	// (Boolean) Sets whether the alert should be paused or not. Defaults to false.
	// Sets whether the alert should be paused or not. Defaults to `false`.
	IsPaused *bool `json:"isPaused,omitempty" tf:"is_paused,omitempty"`

	// value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to map[].
	// Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the rule group.
	// The name of the alert rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Describes what state to enter when the rule's query returns No Data. Options are OK, NoData, and Alerting. Defaults to NoData.
	// Describes what state to enter when the rule's query returns No Data. Options are OK, NoData, and Alerting. Defaults to `NoData`.
	NoDataState *string `json:"noDataState,omitempty" tf:"no_data_state,omitempty"`

	// (String) The unique identifier of the alert rule.
	// The unique identifier of the alert rule.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type RuleParameters struct {

	// value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to map[].
	// Key-value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to `map[]`.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// (String) The ref_id of the query node in the data field to use as the alert condition.
	// The `ref_id` of the query node in the `data` field to use as the alert condition.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition" tf:"condition,omitempty"`

	// (Block List, Min: 1) A sequence of stages that describe the contents of the rule. (see below for nested schema)
	// A sequence of stages that describe the contents of the rule.
	// +kubebuilder:validation:Optional
	Data []DataParameters `json:"data" tf:"data,omitempty"`

	// (String) Describes what state to enter when the rule's query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to Alerting.
	// Describes what state to enter when the rule's query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to `Alerting`.
	// +kubebuilder:validation:Optional
	ExecErrState *string `json:"execErrState,omitempty" tf:"exec_err_state,omitempty"`

	// (String) The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to 0.
	// The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
	// +kubebuilder:validation:Optional
	For *string `json:"for,omitempty" tf:"for,omitempty"`

	// (Boolean) Sets whether the alert should be paused or not. Defaults to false.
	// Sets whether the alert should be paused or not. Defaults to `false`.
	// +kubebuilder:validation:Optional
	IsPaused *bool `json:"isPaused,omitempty" tf:"is_paused,omitempty"`

	// value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to map[].
	// Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the rule group.
	// The name of the alert rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Describes what state to enter when the rule's query returns No Data. Options are OK, NoData, and Alerting. Defaults to NoData.
	// Describes what state to enter when the rule's query returns No Data. Options are OK, NoData, and Alerting. Defaults to `NoData`.
	// +kubebuilder:validation:Optional
	NoDataState *string `json:"noDataState,omitempty" tf:"no_data_state,omitempty"`
}

// RuleGroupSpec defines the desired state of RuleGroup
type RuleGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleGroupParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RuleGroupInitParameters `json:"initProvider,omitempty"`
}

// RuleGroupStatus defines the observed state of RuleGroup.
type RuleGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RuleGroup is the Schema for the RuleGroups API. Manages Grafana Alerting rule groups. Official documentation https://grafana.com/docs/grafana/latest/alerting/alerting-rules/HTTP API https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#alert-rules This resource requires Grafana 9.1.0 or later.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type RuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.intervalSeconds) || (has(self.initProvider) && has(self.initProvider.intervalSeconds))",message="spec.forProvider.intervalSeconds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule) || (has(self.initProvider) && has(self.initProvider.rule))",message="spec.forProvider.rule is a required parameter"
	Spec   RuleGroupSpec   `json:"spec"`
	Status RuleGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleGroupList contains a list of RuleGroups
type RuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleGroup `json:"items"`
}

// Repository type metadata.
var (
	RuleGroup_Kind             = "RuleGroup"
	RuleGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleGroup_Kind}.String()
	RuleGroup_KindAPIVersion   = RuleGroup_Kind + "." + CRDGroupVersion.String()
	RuleGroup_GroupVersionKind = CRDGroupVersion.WithKind(RuleGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleGroup{}, &RuleGroupList{})
}
