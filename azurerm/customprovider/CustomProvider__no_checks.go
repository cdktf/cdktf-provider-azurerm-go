// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package customprovider

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomProvider) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validatePutActionParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validatePutResourceTypeParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validatePutTimeoutsParameters(value *CustomProviderTimeouts) error {
	return nil
}

func (c *jsiiProxy_CustomProvider) validatePutValidationParameters(value interface{}) error {
	return nil
}

func validateCustomProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCustomProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCustomProvider_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomProvider) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func validateNewCustomProviderParameters(scope constructs.Construct, id *string, config *CustomProviderConfig) error {
	return nil
}

