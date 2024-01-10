// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package keyvault

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyVaultAccessPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KeyVaultAccessPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KeyVaultAccessPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KeyVaultAccessPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVaultAccessPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultAccessPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KeyVaultAccessPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKeyVaultAccessPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

