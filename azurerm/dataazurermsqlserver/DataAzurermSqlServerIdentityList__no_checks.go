// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermsqlserver

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermSqlServerIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermSqlServerIdentityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermSqlServerIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermSqlServerIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermSqlServerIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermSqlServerIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

