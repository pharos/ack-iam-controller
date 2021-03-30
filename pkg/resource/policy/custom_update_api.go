package policy

import (
	"context"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go/service/iam"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (rm *resourceManager) customUpdatePolicy(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {

	input, err := rm.newCreatePolicyVersionPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	policyErr := rm.prunePolicyVersions(ctx, *input.PolicyArn)
	if policyErr != nil {
		return nil, policyErr
	}

	if err != nil {
		return nil, err
	}
	resp, respErr := rm.sdkapi.CreatePolicyVersionWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}

	ko := latest.ko.DeepCopy()

	if resp.PolicyVersion.CreateDate != nil {
		ko.Status.CreateDate = &metav1.Time{*resp.PolicyVersion.CreateDate}
	}
	if resp.PolicyVersion.VersionId != nil {
		ko.Status.DefaultVersionID = resp.PolicyVersion.VersionId
	}
	if resp.PolicyVersion.Document != nil {
		ko.Spec.PolicyDocument = resp.PolicyVersion.Document
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreatePolicyVersionPayload returns an SDK-specific struct for the HTTP request
// payload of the CreatePolicyVersion API call for the resource
func (rm *resourceManager) newCreatePolicyVersionPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreatePolicyVersionInput, error) {
	res := &svcsdk.CreatePolicyVersionInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetPolicyArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	} else {
		res.SetPolicyArn(rm.ARNFromName(*r.ko.Spec.PolicyName))
	}

	if r.ko.Spec.PolicyDocument != nil {
		res.SetPolicyDocument(*r.ko.Spec.PolicyDocument)
	}
	res.SetSetAsDefault(true)

	return res, nil
}

// prunePolicyVersions deletes the oldest versions.
//
// Old versions are deleted until there are 4 or less remaining, which means at
// least one more can be created before hitting the maximum of 5.
//
// The default version is never deleted.
func (rm *resourceManager) prunePolicyVersions(
	ctx context.Context,
	arn string,
) error {
	versions, err := rm.listPolicyVersions(ctx, arn)
	if err != nil {
		return err
	}
	if len(versions) < 5 {
		return nil
	}

	var oldestVersion *svcsdk.PolicyVersion
	for _, version := range versions {
		if *version.IsDefaultVersion {
			continue
		}
		if oldestVersion == nil ||
			version.CreateDate.Before(*oldestVersion.CreateDate) {
			oldestVersion = version
		}
	}

	err1 := rm.deletePolicyVersion(ctx, &arn, oldestVersion)
	return err1
}

func (rm *resourceManager) listPolicyVersions(
	ctx context.Context,
	arn string,
) ([]*svcsdk.PolicyVersion, error) {

	request := &svcsdk.ListPolicyVersionsInput{}
	request.SetPolicyArn(arn)

	response, err := rm.sdkapi.ListPolicyVersionsWithContext(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Versions, nil
}

func (rm *resourceManager) deletePolicyVersion(
	ctx context.Context,
	arn *string,
	policyVersion *svcsdk.PolicyVersion,
) error {
	request := &svcsdk.DeletePolicyVersionInput{
		PolicyArn: arn,
		VersionId: policyVersion.VersionId,
	}

	_, err := rm.sdkapi.DeletePolicyVersionWithContext(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
