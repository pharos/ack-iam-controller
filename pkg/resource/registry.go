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

package resource

import (
	ackrt "github.com/aws-controllers-k8s/runtime/pkg/runtime"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"
)

var (
	reg = ackrt.NewRegistry()
)

// GetManagerFactories returns a slice of resource manager factories that are
// registered with this package
func GetManagerFactories() []acktypes.AWSResourceManagerFactory {
	return reg.GetResourceManagerFactories()
}

// RegisterManagerFactory registers a resource manager factory with the
// package's registry
func RegisterManagerFactory(f acktypes.AWSResourceManagerFactory) {
	reg.RegisterResourceManagerFactory(f)
}
