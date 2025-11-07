// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermloadtest

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermLoadTestEncryptionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzurermLoadTestEncryptionList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermLoadTestEncryptionList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermLoadTestEncryptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermLoadTestEncryptionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermLoadTestEncryptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermLoadTestEncryptionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

