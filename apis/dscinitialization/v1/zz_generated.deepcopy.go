//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneSpec) DeepCopyInto(out *ControlPlaneSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneSpec.
func (in *ControlPlaneSpec) DeepCopy() *ControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSCInitialization) DeepCopyInto(out *DSCInitialization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSCInitialization.
func (in *DSCInitialization) DeepCopy() *DSCInitialization {
	if in == nil {
		return nil
	}
	out := new(DSCInitialization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DSCInitialization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSCInitializationList) DeepCopyInto(out *DSCInitializationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DSCInitialization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSCInitializationList.
func (in *DSCInitializationList) DeepCopy() *DSCInitializationList {
	if in == nil {
		return nil
	}
	out := new(DSCInitializationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DSCInitializationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSCInitializationSpec) DeepCopyInto(out *DSCInitializationSpec) {
	*out = *in
	out.Monitoring = in.Monitoring
	out.ServiceMesh = in.ServiceMesh
	out.Serverless = in.Serverless
	out.DevFlags = in.DevFlags
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSCInitializationSpec.
func (in *DSCInitializationSpec) DeepCopy() *DSCInitializationSpec {
	if in == nil {
		return nil
	}
	out := new(DSCInitializationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSCInitializationStatus) DeepCopyInto(out *DSCInitializationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditionsv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSCInitializationStatus.
func (in *DSCInitializationStatus) DeepCopy() *DSCInitializationStatus {
	if in == nil {
		return nil
	}
	out := new(DSCInitializationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevFlags) DeepCopyInto(out *DevFlags) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevFlags.
func (in *DevFlags) DeepCopy() *DevFlags {
	if in == nil {
		return nil
	}
	out := new(DevFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureTracker) DeepCopyInto(out *FeatureTracker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureTracker.
func (in *FeatureTracker) DeepCopy() *FeatureTracker {
	if in == nil {
		return nil
	}
	out := new(FeatureTracker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureTracker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureTrackerList) DeepCopyInto(out *FeatureTrackerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FeatureTracker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureTrackerList.
func (in *FeatureTrackerList) DeepCopy() *FeatureTrackerList {
	if in == nil {
		return nil
	}
	out := new(FeatureTrackerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureTrackerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureTrackerSpec) DeepCopyInto(out *FeatureTrackerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureTrackerSpec.
func (in *FeatureTrackerSpec) DeepCopy() *FeatureTrackerSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureTrackerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureTrackerStatus) DeepCopyInto(out *FeatureTrackerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureTrackerStatus.
func (in *FeatureTrackerStatus) DeepCopy() *FeatureTrackerStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureTrackerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressGatewaySpec) DeepCopyInto(out *IngressGatewaySpec) {
	*out = *in
	out.Certificate = in.Certificate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressGatewaySpec.
func (in *IngressGatewaySpec) DeepCopy() *IngressGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(IngressGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessSpec) DeepCopyInto(out *ServerlessSpec) {
	*out = *in
	out.Serving = in.Serving
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessSpec.
func (in *ServerlessSpec) DeepCopy() *ServerlessSpec {
	if in == nil {
		return nil
	}
	out := new(ServerlessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshSpec) DeepCopyInto(out *ServiceMeshSpec) {
	*out = *in
	out.ControlPlane = in.ControlPlane
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshSpec.
func (in *ServiceMeshSpec) DeepCopy() *ServiceMeshSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingSpec) DeepCopyInto(out *ServingSpec) {
	*out = *in
	out.IngressGateway = in.IngressGateway
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingSpec.
func (in *ServingSpec) DeepCopy() *ServingSpec {
	if in == nil {
		return nil
	}
	out := new(ServingSpec)
	in.DeepCopyInto(out)
	return out
}
