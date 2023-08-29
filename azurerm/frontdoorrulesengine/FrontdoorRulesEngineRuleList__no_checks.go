// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package frontdoorrulesengine

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FrontdoorRulesEngineRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FrontdoorRulesEngineRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FrontdoorRulesEngineRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FrontdoorRulesEngineRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FrontdoorRulesEngineRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FrontdoorRulesEngineRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFrontdoorRulesEngineRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

