// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermimage

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermImageDataDiskList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzurermImageDataDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermImageDataDiskList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImageDataDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImageDataDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImageDataDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermImageDataDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

