// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package keyvaultsecret

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyVaultSecret) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateImportFromParameters(id *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateMoveToIdParameters(id *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultSecret) validatePutTimeoutsParameters(value *KeyVaultSecretTimeouts) error {
	return nil
}

func validateKeyVaultSecret_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateKeyVaultSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKeyVaultSecret_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateKeyVaultSecret_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetContentTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetExpirationDateParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetKeyVaultIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetNotBeforeDateParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultSecret) validateSetValueParameters(val *string) error {
	return nil
}

func validateNewKeyVaultSecretParameters(scope constructs.Construct, id *string, config *KeyVaultSecretConfig) error {
	return nil
}

