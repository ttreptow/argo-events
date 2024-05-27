//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 BlackRock, Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	common "github.com/argoproj/argo-events/pkg/apis/common"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BusConfig) DeepCopyInto(out *BusConfig) {
	*out = *in
	if in.NATS != nil {
		in, out := &in.NATS, &out.NATS
		*out = new(NATSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.JetStream != nil {
		in, out := &in.JetStream, &out.JetStream
		*out = new(JetStreamConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = new(KafkaBus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BusConfig.
func (in *BusConfig) DeepCopy() *BusConfig {
	if in == nil {
		return nil
	}
	out := new(BusConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerTemplate) DeepCopyInto(out *ContainerTemplate) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerTemplate.
func (in *ContainerTemplate) DeepCopy() *ContainerTemplate {
	if in == nil {
		return nil
	}
	out := new(ContainerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBus) DeepCopyInto(out *EventBus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBus.
func (in *EventBus) DeepCopy() *EventBus {
	if in == nil {
		return nil
	}
	out := new(EventBus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventBus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBusList) DeepCopyInto(out *EventBusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventBus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBusList.
func (in *EventBusList) DeepCopy() *EventBusList {
	if in == nil {
		return nil
	}
	out := new(EventBusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventBusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBusSpec) DeepCopyInto(out *EventBusSpec) {
	*out = *in
	if in.NATS != nil {
		in, out := &in.NATS, &out.NATS
		*out = new(NATSBus)
		(*in).DeepCopyInto(*out)
	}
	if in.JetStream != nil {
		in, out := &in.JetStream, &out.JetStream
		*out = new(JetStreamBus)
		(*in).DeepCopyInto(*out)
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = new(KafkaBus)
		(*in).DeepCopyInto(*out)
	}
	if in.JetStreamExotic != nil {
		in, out := &in.JetStreamExotic, &out.JetStreamExotic
		*out = new(JetStreamConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBusSpec.
func (in *EventBusSpec) DeepCopy() *EventBusSpec {
	if in == nil {
		return nil
	}
	out := new(EventBusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBusStatus) DeepCopyInto(out *EventBusStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBusStatus.
func (in *EventBusStatus) DeepCopy() *EventBusStatus {
	if in == nil {
		return nil
	}
	out := new(EventBusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JetStreamBus) DeepCopyInto(out *JetStreamBus) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ContainerTemplate != nil {
		in, out := &in.ContainerTemplate, &out.ContainerTemplate
		*out = new(ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.ReloaderContainerTemplate != nil {
		in, out := &in.ReloaderContainerTemplate, &out.ReloaderContainerTemplate
		*out = new(ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricsContainerTemplate != nil {
		in, out := &in.MetricsContainerTemplate, &out.MetricsContainerTemplate
		*out = new(ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(common.Metadata)
		(*in).DeepCopyInto(*out)
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
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int32)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(string)
		**out = **in
	}
	if in.StartArgs != nil {
		in, out := &in.StartArgs, &out.StartArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StreamConfig != nil {
		in, out := &in.StreamConfig, &out.StreamConfig
		*out = new(string)
		**out = **in
	}
	if in.MaxPayload != nil {
		in, out := &in.MaxPayload, &out.MaxPayload
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JetStreamBus.
func (in *JetStreamBus) DeepCopy() *JetStreamBus {
	if in == nil {
		return nil
	}
	out := new(JetStreamBus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JetStreamConfig) DeepCopyInto(out *JetStreamConfig) {
	*out = *in
	if in.AccessSecret != nil {
		in, out := &in.AccessSecret, &out.AccessSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	out.TLS = in.TLS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JetStreamConfig.
func (in *JetStreamConfig) DeepCopy() *JetStreamConfig {
	if in == nil {
		return nil
	}
	out := new(JetStreamConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JetStreamTLSConfig) DeepCopyInto(out *JetStreamTLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JetStreamTLSConfig.
func (in *JetStreamTLSConfig) DeepCopy() *JetStreamTLSConfig {
	if in == nil {
		return nil
	}
	out := new(JetStreamTLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaBus) DeepCopyInto(out *KafkaBus) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(common.TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SASL != nil {
		in, out := &in.SASL, &out.SASL
		*out = new(common.SASLConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ConsumerGroup != nil {
		in, out := &in.ConsumerGroup, &out.ConsumerGroup
		*out = new(KafkaConsumerGroup)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaBus.
func (in *KafkaBus) DeepCopy() *KafkaBus {
	if in == nil {
		return nil
	}
	out := new(KafkaBus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConsumerGroup) DeepCopyInto(out *KafkaConsumerGroup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConsumerGroup.
func (in *KafkaConsumerGroup) DeepCopy() *KafkaConsumerGroup {
	if in == nil {
		return nil
	}
	out := new(KafkaConsumerGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSBus) DeepCopyInto(out *NATSBus) {
	*out = *in
	if in.Native != nil {
		in, out := &in.Native, &out.Native
		*out = new(NativeStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Exotic != nil {
		in, out := &in.Exotic, &out.Exotic
		*out = new(NATSConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSBus.
func (in *NATSBus) DeepCopy() *NATSBus {
	if in == nil {
		return nil
	}
	out := new(NATSBus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSConfig) DeepCopyInto(out *NATSConfig) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(AuthStrategy)
		**out = **in
	}
	if in.AccessSecret != nil {
		in, out := &in.AccessSecret, &out.AccessSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSConfig.
func (in *NATSConfig) DeepCopy() *NATSConfig {
	if in == nil {
		return nil
	}
	out := new(NATSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NativeStrategy) DeepCopyInto(out *NativeStrategy) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(AuthStrategy)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerTemplate != nil {
		in, out := &in.ContainerTemplate, &out.ContainerTemplate
		*out = new(ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricsContainerTemplate != nil {
		in, out := &in.MetricsContainerTemplate, &out.MetricsContainerTemplate
		*out = new(ContainerTemplate)
		(*in).DeepCopyInto(*out)
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
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(common.Metadata)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxAge != nil {
		in, out := &in.MaxAge, &out.MaxAge
		*out = new(string)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int32)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxMsgs != nil {
		in, out := &in.MaxMsgs, &out.MaxMsgs
		*out = new(uint64)
		**out = **in
	}
	if in.MaxBytes != nil {
		in, out := &in.MaxBytes, &out.MaxBytes
		*out = new(string)
		**out = **in
	}
	if in.MaxSubs != nil {
		in, out := &in.MaxSubs, &out.MaxSubs
		*out = new(uint64)
		**out = **in
	}
	if in.MaxPayload != nil {
		in, out := &in.MaxPayload, &out.MaxPayload
		*out = new(string)
		**out = **in
	}
	if in.RaftHeartbeatTimeout != nil {
		in, out := &in.RaftHeartbeatTimeout, &out.RaftHeartbeatTimeout
		*out = new(string)
		**out = **in
	}
	if in.RaftElectionTimeout != nil {
		in, out := &in.RaftElectionTimeout, &out.RaftElectionTimeout
		*out = new(string)
		**out = **in
	}
	if in.RaftLeaseTimeout != nil {
		in, out := &in.RaftLeaseTimeout, &out.RaftLeaseTimeout
		*out = new(string)
		**out = **in
	}
	if in.RaftCommitTimeout != nil {
		in, out := &in.RaftCommitTimeout, &out.RaftCommitTimeout
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NativeStrategy.
func (in *NativeStrategy) DeepCopy() *NativeStrategy {
	if in == nil {
		return nil
	}
	out := new(NativeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceStrategy) DeepCopyInto(out *PersistenceStrategy) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(v1.PersistentVolumeAccessMode)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceStrategy.
func (in *PersistenceStrategy) DeepCopy() *PersistenceStrategy {
	if in == nil {
		return nil
	}
	out := new(PersistenceStrategy)
	in.DeepCopyInto(out)
	return out
}
