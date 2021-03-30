package role_policy

import (
	"context"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

func (rm *resourceManager) customUpdateRolePolicy(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {

	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	_, respErr := rm.sdkapi.PutRolePolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "PutRolePolicy", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}
