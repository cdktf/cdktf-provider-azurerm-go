// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package healthcareworkspace

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthcareWorkspacePrivateEndpointConnectionList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HealthcareWorkspacePrivateEndpointConnectionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HealthcareWorkspacePrivateEndpointConnectionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthcareWorkspacePrivateEndpointConnectionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HealthcareWorkspacePrivateEndpointConnectionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHealthcareWorkspacePrivateEndpointConnectionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

