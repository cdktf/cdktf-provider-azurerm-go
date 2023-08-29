// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mobilenetworkservice

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MobileNetworkServicePccRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MobileNetworkServicePccRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MobileNetworkServicePccRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MobileNetworkServicePccRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MobileNetworkServicePccRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MobileNetworkServicePccRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMobileNetworkServicePccRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

