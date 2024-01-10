// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package windowswebapp

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsWebAppStorageAccountList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WindowsWebAppStorageAccountList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WindowsWebAppStorageAccountList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WindowsWebAppStorageAccountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWindowsWebAppStorageAccountListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

