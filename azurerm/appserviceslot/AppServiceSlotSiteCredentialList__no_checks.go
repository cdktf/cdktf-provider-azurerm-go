// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appserviceslot

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppServiceSlotSiteCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppServiceSlotSiteCredentialListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

