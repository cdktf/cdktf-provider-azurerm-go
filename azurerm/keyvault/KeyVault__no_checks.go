// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package keyvault

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyVault) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateImportFromParameters(id *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validatePutAccessPolicyParameters(value interface{}) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validatePutContactParameters(value interface{}) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validatePutNetworkAclsParameters(value *KeyVaultNetworkAcls) error {
	return nil
}

func (k *jsiiProxy_KeyVault) validatePutTimeoutsParameters(value *KeyVaultTimeouts) error {
	return nil
}

func validateKeyVault_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateKeyVault_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKeyVault_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateKeyVault_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetEnabledForDeploymentParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetEnabledForDiskEncryptionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetEnabledForTemplateDeploymentParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetEnableRbacAuthorizationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetPublicNetworkAccessEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetPurgeProtectionEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetSkuNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetSoftDeleteRetentionDaysParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_KeyVault) validateSetTenantIdParameters(val *string) error {
	return nil
}

func validateNewKeyVaultParameters(scope constructs.Construct, id *string, config *KeyVaultConfig) error {
	return nil
}

