//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKey) DeepCopyInto(out *APIKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKey.
func (in *APIKey) DeepCopy() *APIKey {
	if in == nil {
		return nil
	}
	out := new(APIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyList) DeepCopyInto(out *APIKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyList.
func (in *APIKeyList) DeepCopy() *APIKeyList {
	if in == nil {
		return nil
	}
	out := new(APIKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyObservation) DeepCopyInto(out *APIKeyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyObservation.
func (in *APIKeyObservation) DeepCopy() *APIKeyObservation {
	if in == nil {
		return nil
	}
	out := new(APIKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyParameters) DeepCopyInto(out *APIKeyParameters) {
	*out = *in
	if in.CloudOrgSlug != nil {
		in, out := &in.CloudOrgSlug, &out.CloudOrgSlug
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyParameters.
func (in *APIKeyParameters) DeepCopy() *APIKeyParameters {
	if in == nil {
		return nil
	}
	out := new(APIKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeySpec) DeepCopyInto(out *APIKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeySpec.
func (in *APIKeySpec) DeepCopy() *APIKeySpec {
	if in == nil {
		return nil
	}
	out := new(APIKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyStatus) DeepCopyInto(out *APIKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyStatus.
func (in *APIKeyStatus) DeepCopy() *APIKeyStatus {
	if in == nil {
		return nil
	}
	out := new(APIKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stack) DeepCopyInto(out *Stack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stack.
func (in *Stack) DeepCopy() *Stack {
	if in == nil {
		return nil
	}
	out := new(Stack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackList) DeepCopyInto(out *StackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackList.
func (in *StackList) DeepCopy() *StackList {
	if in == nil {
		return nil
	}
	out := new(StackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackObservation) DeepCopyInto(out *StackObservation) {
	*out = *in
	if in.AlertmanagerName != nil {
		in, out := &in.AlertmanagerName, &out.AlertmanagerName
		*out = new(string)
		**out = **in
	}
	if in.AlertmanagerStatus != nil {
		in, out := &in.AlertmanagerStatus, &out.AlertmanagerStatus
		*out = new(string)
		**out = **in
	}
	if in.AlertmanagerURL != nil {
		in, out := &in.AlertmanagerURL, &out.AlertmanagerURL
		*out = new(string)
		**out = **in
	}
	if in.AlertmanagerUserID != nil {
		in, out := &in.AlertmanagerUserID, &out.AlertmanagerUserID
		*out = new(float64)
		**out = **in
	}
	if in.GraphiteName != nil {
		in, out := &in.GraphiteName, &out.GraphiteName
		*out = new(string)
		**out = **in
	}
	if in.GraphiteStatus != nil {
		in, out := &in.GraphiteStatus, &out.GraphiteStatus
		*out = new(string)
		**out = **in
	}
	if in.GraphiteURL != nil {
		in, out := &in.GraphiteURL, &out.GraphiteURL
		*out = new(string)
		**out = **in
	}
	if in.GraphiteUserID != nil {
		in, out := &in.GraphiteUserID, &out.GraphiteUserID
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LogsName != nil {
		in, out := &in.LogsName, &out.LogsName
		*out = new(string)
		**out = **in
	}
	if in.LogsStatus != nil {
		in, out := &in.LogsStatus, &out.LogsStatus
		*out = new(string)
		**out = **in
	}
	if in.LogsURL != nil {
		in, out := &in.LogsURL, &out.LogsURL
		*out = new(string)
		**out = **in
	}
	if in.LogsUserID != nil {
		in, out := &in.LogsUserID, &out.LogsUserID
		*out = new(float64)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(float64)
		**out = **in
	}
	if in.OrgName != nil {
		in, out := &in.OrgName, &out.OrgName
		*out = new(string)
		**out = **in
	}
	if in.OrgSlug != nil {
		in, out := &in.OrgSlug, &out.OrgSlug
		*out = new(string)
		**out = **in
	}
	if in.PrometheusName != nil {
		in, out := &in.PrometheusName, &out.PrometheusName
		*out = new(string)
		**out = **in
	}
	if in.PrometheusRemoteEndpoint != nil {
		in, out := &in.PrometheusRemoteEndpoint, &out.PrometheusRemoteEndpoint
		*out = new(string)
		**out = **in
	}
	if in.PrometheusRemoteWriteEndpoint != nil {
		in, out := &in.PrometheusRemoteWriteEndpoint, &out.PrometheusRemoteWriteEndpoint
		*out = new(string)
		**out = **in
	}
	if in.PrometheusStatus != nil {
		in, out := &in.PrometheusStatus, &out.PrometheusStatus
		*out = new(string)
		**out = **in
	}
	if in.PrometheusURL != nil {
		in, out := &in.PrometheusURL, &out.PrometheusURL
		*out = new(string)
		**out = **in
	}
	if in.PrometheusUserID != nil {
		in, out := &in.PrometheusUserID, &out.PrometheusUserID
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TracesName != nil {
		in, out := &in.TracesName, &out.TracesName
		*out = new(string)
		**out = **in
	}
	if in.TracesStatus != nil {
		in, out := &in.TracesStatus, &out.TracesStatus
		*out = new(string)
		**out = **in
	}
	if in.TracesURL != nil {
		in, out := &in.TracesURL, &out.TracesURL
		*out = new(string)
		**out = **in
	}
	if in.TracesUserID != nil {
		in, out := &in.TracesUserID, &out.TracesUserID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackObservation.
func (in *StackObservation) DeepCopy() *StackObservation {
	if in == nil {
		return nil
	}
	out := new(StackObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackParameters) DeepCopyInto(out *StackParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RegionSlug != nil {
		in, out := &in.RegionSlug, &out.RegionSlug
		*out = new(string)
		**out = **in
	}
	if in.Slug != nil {
		in, out := &in.Slug, &out.Slug
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.WaitForReadiness != nil {
		in, out := &in.WaitForReadiness, &out.WaitForReadiness
		*out = new(bool)
		**out = **in
	}
	if in.WaitForReadinessTimeout != nil {
		in, out := &in.WaitForReadinessTimeout, &out.WaitForReadinessTimeout
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackParameters.
func (in *StackParameters) DeepCopy() *StackParameters {
	if in == nil {
		return nil
	}
	out := new(StackParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSpec) DeepCopyInto(out *StackSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSpec.
func (in *StackSpec) DeepCopy() *StackSpec {
	if in == nil {
		return nil
	}
	out := new(StackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackStatus) DeepCopyInto(out *StackStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackStatus.
func (in *StackStatus) DeepCopy() *StackStatus {
	if in == nil {
		return nil
	}
	out := new(StackStatus)
	in.DeepCopyInto(out)
	return out
}