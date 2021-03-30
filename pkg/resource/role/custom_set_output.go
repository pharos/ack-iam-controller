package role

import (
	"context"

	svcapitypes "github.com/aws-controllers-k8s/iam-controller/apis/v1alpha1"
	svcsdk "github.com/aws/aws-sdk-go/service/iam"
)

func (rm *resourceManager) CustomGetRoleSetOutput(
	ctx context.Context,
	r *resource,
	resp *svcsdk.GetRoleOutput,
	ko *svcapitypes.Role,
) (*svcapitypes.Role, error) {

	if resp.Role.PermissionsBoundary != nil {
		ko.Spec.PermissionsBoundary = resp.Role.PermissionsBoundary.PermissionsBoundaryArn
	}

	return ko, nil
}
