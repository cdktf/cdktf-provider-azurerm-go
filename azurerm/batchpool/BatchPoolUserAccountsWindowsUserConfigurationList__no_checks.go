// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package batchpool

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchPoolUserAccountsWindowsUserConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchPoolUserAccountsWindowsUserConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

