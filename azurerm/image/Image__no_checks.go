// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package image

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Image) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateImportFromParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (i *jsiiProxy_Image) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (i *jsiiProxy_Image) validatePutDataDiskParameters(value interface{}) error {
	return nil
}

func (i *jsiiProxy_Image) validatePutOsDiskParameters(value *ImageOsDisk) error {
	return nil
}

func (i *jsiiProxy_Image) validatePutTimeoutsParameters(value *ImageTimeouts) error {
	return nil
}

func validateImage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateImage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateImage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateImage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetHyperVGenerationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetSourceVirtualMachineIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Image) validateSetZoneResilientParameters(val interface{}) error {
	return nil
}

func validateNewImageParameters(scope constructs.Construct, id *string, config *ImageConfig) error {
	return nil
}

