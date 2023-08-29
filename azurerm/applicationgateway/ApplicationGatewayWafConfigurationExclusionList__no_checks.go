// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package applicationgateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationGatewayWafConfigurationExclusionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationExclusionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationExclusionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationExclusionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationExclusionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationExclusionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationGatewayWafConfigurationExclusionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

