// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RolePolicyAttachmentSpec defines the desired state of RolePolicyAttachment
type RolePolicyAttachmentSpec struct {
	// The Amazon Resource Name (ARN) of the IAM policy you want to attach.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	// +kubebuilder:validation:Required
	PolicyARN *string `json:"policyARN"`
	// The name (friendly name, not ARN) of the role to attach the policy to.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	// +kubebuilder:validation:Required
	RoleName *string `json:"roleName"`
}

// RolePolicyAttachmentStatus defines the observed state of RolePolicyAttachment
type RolePolicyAttachmentStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// RolePolicyAttachment is the Schema for the RolePolicyAttachments API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type RolePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RolePolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            RolePolicyAttachmentStatus `json:"status,omitempty"`
}

// RolePolicyAttachmentList contains a list of RolePolicyAttachment
// +kubebuilder:object:root=true
type RolePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RolePolicyAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RolePolicyAttachment{}, &RolePolicyAttachmentList{})
}
