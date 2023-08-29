// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermmssqlserver

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermMssqlServerIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermMssqlServerIdentityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermMssqlServerIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermMssqlServerIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermMssqlServerIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermMssqlServerIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

