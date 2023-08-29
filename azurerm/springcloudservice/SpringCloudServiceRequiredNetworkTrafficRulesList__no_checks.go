// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package springcloudservice

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpringCloudServiceRequiredNetworkTrafficRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpringCloudServiceRequiredNetworkTrafficRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

