// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package keyvault

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyVaultContactList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultContactList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KeyVaultContactList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KeyVaultContactList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVaultContactList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultContactList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KeyVaultContactList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKeyVaultContactListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

