//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationFilter) DeepCopyInto(out *AuthenticationFilter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationFilter.
func (in *AuthenticationFilter) DeepCopy() *AuthenticationFilter {
	if in == nil {
		return nil
	}
	out := new(AuthenticationFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthenticationFilter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationFilterList) DeepCopyInto(out *AuthenticationFilterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthenticationFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationFilterList.
func (in *AuthenticationFilterList) DeepCopy() *AuthenticationFilterList {
	if in == nil {
		return nil
	}
	out := new(AuthenticationFilterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthenticationFilterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationFilterSpec) DeepCopyInto(out *AuthenticationFilterSpec) {
	*out = *in
	if in.JwtProviders != nil {
		in, out := &in.JwtProviders, &out.JwtProviders
		*out = make([]JwtAuthenticationFilterProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationFilterSpec.
func (in *AuthenticationFilterSpec) DeepCopy() *AuthenticationFilterSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationFilterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalRateLimit) DeepCopyInto(out *GlobalRateLimit) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RateLimitRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalRateLimit.
func (in *GlobalRateLimit) DeepCopy() *GlobalRateLimit {
	if in == nil {
		return nil
	}
	out := new(GlobalRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderMatch) DeepCopyInto(out *HeaderMatch) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(HeaderMatchType)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatch.
func (in *HeaderMatch) DeepCopy() *HeaderMatch {
	if in == nil {
		return nil
	}
	out := new(HeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JwtAuthenticationFilterProvider) DeepCopyInto(out *JwtAuthenticationFilterProvider) {
	*out = *in
	if in.Audiences != nil {
		in, out := &in.Audiences, &out.Audiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.RemoteJWKS = in.RemoteJWKS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JwtAuthenticationFilterProvider.
func (in *JwtAuthenticationFilterProvider) DeepCopy() *JwtAuthenticationFilterProvider {
	if in == nil {
		return nil
	}
	out := new(JwtAuthenticationFilterProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitFilter) DeepCopyInto(out *RateLimitFilter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitFilter.
func (in *RateLimitFilter) DeepCopy() *RateLimitFilter {
	if in == nil {
		return nil
	}
	out := new(RateLimitFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitFilter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitFilterSpec) DeepCopyInto(out *RateLimitFilterSpec) {
	*out = *in
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(GlobalRateLimit)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitFilterSpec.
func (in *RateLimitFilterSpec) DeepCopy() *RateLimitFilterSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitFilterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitList) DeepCopyInto(out *RateLimitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimitFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitList.
func (in *RateLimitList) DeepCopy() *RateLimitList {
	if in == nil {
		return nil
	}
	out := new(RateLimitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitRule) DeepCopyInto(out *RateLimitRule) {
	*out = *in
	if in.ClientSelectors != nil {
		in, out := &in.ClientSelectors, &out.ClientSelectors
		*out = make([]RateLimitSelectCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Limit = in.Limit
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitRule.
func (in *RateLimitRule) DeepCopy() *RateLimitRule {
	if in == nil {
		return nil
	}
	out := new(RateLimitRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitSelectCondition) DeepCopyInto(out *RateLimitSelectCondition) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitSelectCondition.
func (in *RateLimitSelectCondition) DeepCopy() *RateLimitSelectCondition {
	if in == nil {
		return nil
	}
	out := new(RateLimitSelectCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitValue) DeepCopyInto(out *RateLimitValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitValue.
func (in *RateLimitValue) DeepCopy() *RateLimitValue {
	if in == nil {
		return nil
	}
	out := new(RateLimitValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteJWKS) DeepCopyInto(out *RemoteJWKS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteJWKS.
func (in *RemoteJWKS) DeepCopy() *RemoteJWKS {
	if in == nil {
		return nil
	}
	out := new(RemoteJWKS)
	in.DeepCopyInto(out)
	return out
}
