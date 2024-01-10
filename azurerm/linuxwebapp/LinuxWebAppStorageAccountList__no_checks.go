// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package linuxwebapp

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxWebAppStorageAccountList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LinuxWebAppStorageAccountList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LinuxWebAppStorageAccountList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppStorageAccountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppStorageAccountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppStorageAccountList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppStorageAccountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLinuxWebAppStorageAccountListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

