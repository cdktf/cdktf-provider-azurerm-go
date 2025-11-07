// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package windowsfunctionapp

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsFunctionAppStorageAccountList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WindowsFunctionAppStorageAccountList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WindowsFunctionAppStorageAccountList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppStorageAccountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppStorageAccountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppStorageAccountList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppStorageAccountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWindowsFunctionAppStorageAccountListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

