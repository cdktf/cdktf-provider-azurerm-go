// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appconfigurationfeature

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppConfigurationFeatureTimewindowFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppConfigurationFeatureTimewindowFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationFeatureTimewindowFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationFeatureTimewindowFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationFeatureTimewindowFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationFeatureTimewindowFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppConfigurationFeatureTimewindowFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

