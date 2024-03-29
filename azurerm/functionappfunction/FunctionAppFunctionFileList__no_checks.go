// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionappfunction

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionAppFunctionFileList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionAppFunctionFileList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionAppFunctionFileList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionAppFunctionFileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionAppFunctionFileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionAppFunctionFileList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionAppFunctionFileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionAppFunctionFileListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

