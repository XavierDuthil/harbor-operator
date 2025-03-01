// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
	out.Operator = in.Operator
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentsTLSSpec) DeepCopyInto(out *ComponentsTLSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentsTLSSpec.
func (in *ComponentsTLSSpec) DeepCopy() *ComponentsTLSSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentsTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorStatus) DeepCopyInto(out *OperatorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorStatus.
func (in *OperatorStatus) DeepCopy() *OperatorStatus {
	if in == nil {
		return nil
	}
	out := new(OperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresConnectTimeout) DeepCopyInto(out *PostgresConnectTimeout) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresConnectTimeout.
func (in *PostgresConnectTimeout) DeepCopy() *PostgresConnectTimeout {
	if in == nil {
		return nil
	}
	out := new(PostgresConnectTimeout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresConnection) DeepCopyInto(out *PostgresConnection) {
	*out = *in
	out.PostgresCredentials = in.PostgresCredentials
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]PostgresHostSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresConnection.
func (in *PostgresConnection) DeepCopy() *PostgresConnection {
	if in == nil {
		return nil
	}
	out := new(PostgresConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresConnectionWithParameters) DeepCopyInto(out *PostgresConnectionWithParameters) {
	*out = *in
	in.PostgresConnection.DeepCopyInto(&out.PostgresConnection)
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresConnectionWithParameters.
func (in *PostgresConnectionWithParameters) DeepCopy() *PostgresConnectionWithParameters {
	if in == nil {
		return nil
	}
	out := new(PostgresConnectionWithParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresCredentials) DeepCopyInto(out *PostgresCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresCredentials.
func (in *PostgresCredentials) DeepCopy() *PostgresCredentials {
	if in == nil {
		return nil
	}
	out := new(PostgresCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresHostSpec) DeepCopyInto(out *PostgresHostSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresHostSpec.
func (in *PostgresHostSpec) DeepCopy() *PostgresHostSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresHostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConnection) DeepCopyInto(out *RedisConnection) {
	*out = *in
	out.RedisHostSpec = in.RedisHostSpec
	out.RedisCredentials = in.RedisCredentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConnection.
func (in *RedisConnection) DeepCopy() *RedisConnection {
	if in == nil {
		return nil
	}
	out := new(RedisConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCredentials) DeepCopyInto(out *RedisCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCredentials.
func (in *RedisCredentials) DeepCopy() *RedisCredentials {
	if in == nil {
		return nil
	}
	out := new(RedisCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisHostSpec) DeepCopyInto(out *RedisHostSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisHostSpec.
func (in *RedisHostSpec) DeepCopy() *RedisHostSpec {
	if in == nil {
		return nil
	}
	out := new(RedisHostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrivySeverityTypes) DeepCopyInto(out *TrivySeverityTypes) {
	*out = *in
	if in.Severities != nil {
		in, out := &in.Severities, &out.Severities
		*out = new([]TrivySeverityType)
		if **in != nil {
			in, out := *in, *out
			*out = make([]TrivySeverityType, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrivySeverityTypes.
func (in *TrivySeverityTypes) DeepCopy() *TrivySeverityTypes {
	if in == nil {
		return nil
	}
	out := new(TrivySeverityTypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrivyVulnerabilityTypes) DeepCopyInto(out *TrivyVulnerabilityTypes) {
	*out = *in
	if in.VulnerabilityTypes != nil {
		in, out := &in.VulnerabilityTypes, &out.VulnerabilityTypes
		*out = new([]TrivyVulnerabilityType)
		if **in != nil {
			in, out := *in, *out
			*out = make([]TrivyVulnerabilityType, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrivyVulnerabilityTypes.
func (in *TrivyVulnerabilityTypes) DeepCopy() *TrivyVulnerabilityTypes {
	if in == nil {
		return nil
	}
	out := new(TrivyVulnerabilityTypes)
	in.DeepCopyInto(out)
	return out
}
