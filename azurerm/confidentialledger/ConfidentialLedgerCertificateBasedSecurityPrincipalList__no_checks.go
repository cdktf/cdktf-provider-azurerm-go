// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package confidentialledger

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfidentialLedgerCertificateBasedSecurityPrincipalList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfidentialLedgerCertificateBasedSecurityPrincipalList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfidentialLedgerCertificateBasedSecurityPrincipalList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfidentialLedgerCertificateBasedSecurityPrincipalList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfidentialLedgerCertificateBasedSecurityPrincipalList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfidentialLedgerCertificateBasedSecurityPrincipalList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfidentialLedgerCertificateBasedSecurityPrincipalListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

