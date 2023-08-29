// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apimanagement

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiManagementHostnameConfigurationPortalListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

