// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppServiceSiteConfigIpRestrictionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppServiceSiteConfigIpRestrictionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteConfigIpRestrictionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteConfigIpRestrictionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteConfigIpRestrictionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteConfigIpRestrictionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppServiceSiteConfigIpRestrictionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

