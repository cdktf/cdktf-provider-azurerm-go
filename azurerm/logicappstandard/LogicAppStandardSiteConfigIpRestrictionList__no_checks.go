// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logicappstandard

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogicAppStandardSiteConfigIpRestrictionList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogicAppStandardSiteConfigIpRestrictionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogicAppStandardSiteConfigIpRestrictionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogicAppStandardSiteConfigIpRestrictionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogicAppStandardSiteConfigIpRestrictionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogicAppStandardSiteConfigIpRestrictionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogicAppStandardSiteConfigIpRestrictionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

