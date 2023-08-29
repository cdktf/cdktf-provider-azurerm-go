// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package keyvaultcertificatecontacts

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyVaultCertificateContactsContactList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KeyVaultCertificateContactsContactList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KeyVaultCertificateContactsContactList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyVaultCertificateContactsContactList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyVaultCertificateContactsContactList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KeyVaultCertificateContactsContactList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKeyVaultCertificateContactsContactListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

