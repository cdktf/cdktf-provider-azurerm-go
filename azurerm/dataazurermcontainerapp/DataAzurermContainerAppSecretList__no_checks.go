// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermcontainerapp

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermContainerAppSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermContainerAppSecretList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermContainerAppSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermContainerAppSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermContainerAppSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermContainerAppSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

