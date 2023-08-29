// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logztagrule

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogzTagRuleTagFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogzTagRuleTagFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogzTagRuleTagFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogzTagRuleTagFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogzTagRuleTagFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogzTagRuleTagFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogzTagRuleTagFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

