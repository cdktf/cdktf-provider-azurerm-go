// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermcontainerapp

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermContainerAppTemplateList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermContainerAppTemplateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermContainerAppTemplateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermContainerAppTemplateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermContainerAppTemplateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermContainerAppTemplateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

