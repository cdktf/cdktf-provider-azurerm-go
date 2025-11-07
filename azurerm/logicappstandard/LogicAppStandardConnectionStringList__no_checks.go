// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logicappstandard

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogicAppStandardConnectionStringList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogicAppStandardConnectionStringList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogicAppStandardConnectionStringList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogicAppStandardConnectionStringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogicAppStandardConnectionStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogicAppStandardConnectionStringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogicAppStandardConnectionStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogicAppStandardConnectionStringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

