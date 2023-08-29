// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermimages

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermImagesImagesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermImagesImagesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImagesImagesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImagesImagesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImagesImagesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermImagesImagesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

