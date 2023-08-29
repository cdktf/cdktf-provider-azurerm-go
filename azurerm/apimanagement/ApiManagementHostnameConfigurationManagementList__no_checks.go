// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apimanagement

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiManagementHostnameConfigurationManagementListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

