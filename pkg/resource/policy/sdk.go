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

package policy

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/iam"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/iam-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.IAM{}
	_ = &svcapitypes.Policy{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetPolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetPolicy", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NoSuchEntity" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Policy.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Policy.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Policy.AttachmentCount != nil {
		ko.Status.AttachmentCount = resp.Policy.AttachmentCount
	} else {
		ko.Status.AttachmentCount = nil
	}
	if resp.Policy.CreateDate != nil {
		ko.Status.CreateDate = &metav1.Time{*resp.Policy.CreateDate}
	} else {
		ko.Status.CreateDate = nil
	}
	if resp.Policy.DefaultVersionId != nil {
		ko.Status.DefaultVersionID = resp.Policy.DefaultVersionId
	} else {
		ko.Status.DefaultVersionID = nil
	}
	if resp.Policy.Description != nil {
		ko.Spec.Description = resp.Policy.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.Policy.IsAttachable != nil {
		ko.Status.IsAttachable = resp.Policy.IsAttachable
	} else {
		ko.Status.IsAttachable = nil
	}
	if resp.Policy.Path != nil {
		ko.Spec.Path = resp.Policy.Path
	} else {
		ko.Spec.Path = nil
	}
	if resp.Policy.PermissionsBoundaryUsageCount != nil {
		ko.Status.PermissionsBoundaryUsageCount = resp.Policy.PermissionsBoundaryUsageCount
	} else {
		ko.Status.PermissionsBoundaryUsageCount = nil
	}
	if resp.Policy.PolicyId != nil {
		ko.Status.PolicyID = resp.Policy.PolicyId
	} else {
		ko.Status.PolicyID = nil
	}
	if resp.Policy.PolicyName != nil {
		ko.Spec.PolicyName = resp.Policy.PolicyName
	} else {
		ko.Spec.PolicyName = nil
	}
	if resp.Policy.Tags != nil {
		f10 := []*svcapitypes.Tag{}
		for _, f10iter := range resp.Policy.Tags {
			f10elem := &svcapitypes.Tag{}
			if f10iter.Key != nil {
				f10elem.Key = f10iter.Key
			}
			if f10iter.Value != nil {
				f10elem.Value = f10iter.Value
			}
			f10 = append(f10, f10elem)
		}
		ko.Spec.Tags = f10
	} else {
		ko.Spec.Tags = nil
	}
	if resp.Policy.UpdateDate != nil {
		ko.Status.UpdateDate = &metav1.Time{*resp.Policy.UpdateDate}
	} else {
		ko.Status.UpdateDate = nil
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return (r.ko.Status.ACKResourceMetadata == nil || r.ko.Status.ACKResourceMetadata.ARN == nil)

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetPolicyInput, error) {
	res := &svcsdk.GetPolicyInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetPolicyArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	} else {
		res.SetPolicyArn(rm.ARNFromName(*r.ko.Spec.PolicyName))
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(ctx, r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreatePolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreatePolicy", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Policy.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Policy.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Policy.AttachmentCount != nil {
		ko.Status.AttachmentCount = resp.Policy.AttachmentCount
	} else {
		ko.Status.AttachmentCount = nil
	}
	if resp.Policy.CreateDate != nil {
		ko.Status.CreateDate = &metav1.Time{*resp.Policy.CreateDate}
	} else {
		ko.Status.CreateDate = nil
	}
	if resp.Policy.DefaultVersionId != nil {
		ko.Status.DefaultVersionID = resp.Policy.DefaultVersionId
	} else {
		ko.Status.DefaultVersionID = nil
	}
	if resp.Policy.IsAttachable != nil {
		ko.Status.IsAttachable = resp.Policy.IsAttachable
	} else {
		ko.Status.IsAttachable = nil
	}
	if resp.Policy.PermissionsBoundaryUsageCount != nil {
		ko.Status.PermissionsBoundaryUsageCount = resp.Policy.PermissionsBoundaryUsageCount
	} else {
		ko.Status.PermissionsBoundaryUsageCount = nil
	}
	if resp.Policy.PolicyId != nil {
		ko.Status.PolicyID = resp.Policy.PolicyId
	} else {
		ko.Status.PolicyID = nil
	}
	if resp.Policy.UpdateDate != nil {
		ko.Status.UpdateDate = &metav1.Time{*resp.Policy.UpdateDate}
	} else {
		ko.Status.UpdateDate = nil
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreatePolicyInput, error) {
	res := &svcsdk.CreatePolicyInput{}

	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.Path != nil {
		res.SetPath(*r.ko.Spec.Path)
	}
	if r.ko.Spec.PolicyDocument != nil {
		res.SetPolicyDocument(*r.ko.Spec.PolicyDocument)
	}
	if r.ko.Spec.PolicyName != nil {
		res.SetPolicyName(*r.ko.Spec.PolicyName)
	}
	if r.ko.Spec.Tags != nil {
		f4 := []*svcsdk.Tag{}
		for _, f4iter := range r.ko.Spec.Tags {
			f4elem := &svcsdk.Tag{}
			if f4iter.Key != nil {
				f4elem.SetKey(*f4iter.Key)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTags(f4)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return rm.customUpdatePolicy(ctx, desired, latest, delta)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {

	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeletePolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeletePolicy", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeletePolicyInput, error) {
	res := &svcsdk.DeletePolicyInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetPolicyArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	} else {
		res.SetPolicyArn(rm.ARNFromName(*r.ko.Spec.PolicyName))
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Policy,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Message()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	if terminalCondition != nil || recoverableCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
