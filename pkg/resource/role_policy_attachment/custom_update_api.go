package role_policy_attachment

import (
	"context"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

func (rm *resourceManager) customUpdateRolePolicyAttachment(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {

	// TODO: What if we attach the same managed role again via a different resource object
	//       This should probably be an error, as it's obviously an error in the k8s manifest

	err := rm.sdkDelete(ctx, latest)
	if err != nil {
		return nil, err
	}

	return rm.sdkCreate(ctx, desired)
}
