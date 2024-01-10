// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermappservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermAppServiceConnectionStringList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzurermAppServiceConnectionStringList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermAppServiceConnectionStringList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermAppServiceConnectionStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermAppServiceConnectionStringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermAppServiceConnectionStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermAppServiceConnectionStringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

