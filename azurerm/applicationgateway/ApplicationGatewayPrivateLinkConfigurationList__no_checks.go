// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package applicationgateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationGatewayPrivateLinkConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationGatewayPrivateLinkConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayPrivateLinkConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayPrivateLinkConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayPrivateLinkConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayPrivateLinkConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationGatewayPrivateLinkConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

