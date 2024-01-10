// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppConfigurationPrimaryReadKeyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppConfigurationPrimaryReadKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppConfigurationPrimaryReadKeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryReadKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryReadKeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryReadKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppConfigurationPrimaryReadKeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

