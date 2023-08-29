// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package attestationprovider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttestationProviderPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AttestationProviderPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AttestationProviderPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AttestationProviderPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AttestationProviderPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AttestationProviderPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAttestationProviderPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

