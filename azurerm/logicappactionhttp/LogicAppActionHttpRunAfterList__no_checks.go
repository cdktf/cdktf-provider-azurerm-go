// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logicappactionhttp

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogicAppActionHttpRunAfterList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogicAppActionHttpRunAfterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogicAppActionHttpRunAfterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogicAppActionHttpRunAfterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogicAppActionHttpRunAfterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogicAppActionHttpRunAfterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogicAppActionHttpRunAfterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

