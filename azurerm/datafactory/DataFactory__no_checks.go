// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datafactory

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataFactory) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validatePutGithubConfigurationParameters(value *DataFactoryGithubConfiguration) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validatePutGlobalParameterParameters(value interface{}) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validatePutIdentityParameters(value *DataFactoryIdentity) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validatePutTimeoutsParameters(value *DataFactoryTimeouts) error {
	return nil
}

func (d *jsiiProxy_DataFactory) validatePutVstsConfigurationParameters(value *DataFactoryVstsConfiguration) error {
	return nil
}

func validateDataFactory_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDataFactory_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataFactory_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDataFactory_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetCustomerManagedKeyIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetCustomerManagedKeyIdentityIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetManagedVirtualNetworkEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetPublicNetworkEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetPurviewIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataFactory) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func validateNewDataFactoryParameters(scope constructs.Construct, id *string, config *DataFactoryConfig) error {
	return nil
}

