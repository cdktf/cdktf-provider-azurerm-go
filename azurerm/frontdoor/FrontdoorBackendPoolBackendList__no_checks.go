// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package frontdoor

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FrontdoorBackendPoolBackendList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FrontdoorBackendPoolBackendList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolBackendList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFrontdoorBackendPoolBackendListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

