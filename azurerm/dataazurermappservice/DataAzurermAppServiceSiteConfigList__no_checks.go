// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermappservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermAppServiceSiteConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzurermAppServiceSiteConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermAppServiceSiteConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermAppServiceSiteConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermAppServiceSiteConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermAppServiceSiteConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermAppServiceSiteConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

