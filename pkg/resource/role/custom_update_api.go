package role

import (
	"context"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go/service/iam"
)

// customUpdateRole adds specialized update logic to apply changes to properties not
// included as part of the main UpdateRole API
func (rm *resourceManager) customUpdateRole(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {

	if latest.ko.Spec.AssumeRolePolicyDocument != desired.ko.Spec.AssumeRolePolicyDocument {
		input, err := rm.newUpdateAssumeRolePolicyPayload(ctx, desired)
		if err != nil {
			return nil, err
		}
		_, respErr := rm.sdkapi.UpdateAssumeRolePolicyWithContext(ctx, input)
		if respErr != nil {
			return nil, respErr
		}
	}

	if latest.ko.Spec.PermissionsBoundary != nil && desired.ko.Spec.PermissionsBoundary == nil {
		input, err := rm.newDeleteRolePermissionsBoundaryPayload(ctx, latest)
		if err != nil {
			return nil, err
		}
		_, respErr := rm.sdkapi.DeleteRolePermissionsBoundaryWithContext(ctx, input)
		if respErr != nil {
			return nil, respErr
		}
	} else if latest.ko.Spec.PermissionsBoundary != desired.ko.Spec.PermissionsBoundary {
		input, err := rm.newPutRolePermissionsBoundaryPayload(ctx, desired)
		if err != nil {
			return nil, err
		}
		_, respErr := rm.sdkapi.PutRolePermissionsBoundaryWithContext(ctx, input)
		if respErr != nil {
			return nil, respErr
		}
	}

	return nil, nil
}

// newUpdateAssumeRolePolicyPayload returns an SDK-specific struct for the HTTP request
// payload of the UpdateAssumeRolePolicy API call for the resource
func (rm *resourceManager) newUpdateAssumeRolePolicyPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateAssumeRolePolicyInput, error) {
	res := &svcsdk.UpdateAssumeRolePolicyInput{}

	if r.ko.Spec.AssumeRolePolicyDocument != nil {
		res.SetPolicyDocument(*r.ko.Spec.AssumeRolePolicyDocument)
	}
	if r.ko.Spec.RoleName != nil {
		res.SetRoleName(*r.ko.Spec.RoleName)
	}

	return res, nil
}

// newDeleteRolePermissionsBoundaryPayload returns an SDK-specific struct for the HTTP request
// payload of the DeleteRolePermissionsBoundary API call for the resource
func (rm *resourceManager) newDeleteRolePermissionsBoundaryPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.DeleteRolePermissionsBoundaryInput, error) {
	res := &svcsdk.DeleteRolePermissionsBoundaryInput{}

	if r.ko.Spec.RoleName != nil {
		res.SetRoleName(*r.ko.Spec.RoleName)
	}

	return res, nil
}

// newPutRolePermissionsBoundaryPayload returns an SDK-specific struct for the HTTP request
// payload of the PutRolePermissionsBoundary API call for the resource
func (rm *resourceManager) newPutRolePermissionsBoundaryPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.PutRolePermissionsBoundaryInput, error) {
	res := &svcsdk.PutRolePermissionsBoundaryInput{}

	if r.ko.Spec.PermissionsBoundary != nil {
		res.SetPermissionsBoundary(*r.ko.Spec.PermissionsBoundary)
	}

	if r.ko.Spec.RoleName != nil {
		res.SetRoleName(*r.ko.Spec.RoleName)
	}

	return res, nil
}
