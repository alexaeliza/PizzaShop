// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v1beta2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RaftPartitionState is a state constant for RaftPartition
type RaftPartitionState string

const (
	RaftPartitionInitializing  RaftPartitionState = "Initializing"
	RaftPartitionReconfiguring RaftPartitionState = "Reconfiguring"
	RaftPartitionRunning       RaftPartitionState = "Running"
	RaftPartitionReady         RaftPartitionState = "Ready"
)

type PartitionID uint64

type GroupID uint64

// RaftPartitionSpec specifies a RaftPartitionSpec configuration
type RaftPartitionSpec struct {
	RaftConfig  `json:",inline"`
	Replicas    uint32      `json:"replicas"`
	PartitionID PartitionID `json:"partitionID"`
	GroupID     GroupID     `json:"groupID"`
}

// RaftPartitionStatus defines the status of a RaftPartition
type RaftPartitionStatus struct {
	State     RaftPartitionState            `json:"state,omitempty"`
	Term      *uint64                       `json:"term,omitempty"`
	Leader    *corev1.LocalObjectReference  `json:"leader,omitempty"`
	Followers []corev1.LocalObjectReference `json:"followers,omitempty"`
	Members   uint32                        `json:"members"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RaftPartition is the Schema for the RaftPartition API
// +k8s:openapi-gen=true
type RaftPartition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RaftPartitionSpec   `json:"spec,omitempty"`
	Status            RaftPartitionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RaftPartitionList contains a list of RaftPartition
type RaftPartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the RaftPartition of items in the list
	Items []RaftPartition `json:"items"`
}
