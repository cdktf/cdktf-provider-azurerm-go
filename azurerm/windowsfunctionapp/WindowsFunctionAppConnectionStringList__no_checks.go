// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package windowsfunctionapp

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsFunctionAppConnectionStringList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WindowsFunctionAppConnectionStringList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WindowsFunctionAppConnectionStringList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppConnectionStringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppConnectionStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppConnectionStringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppConnectionStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWindowsFunctionAppConnectionStringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

