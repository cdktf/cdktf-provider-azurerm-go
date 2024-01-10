// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package healthcaredicomservice

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthcareDicomServiceAuthenticationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HealthcareDicomServiceAuthenticationList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HealthcareDicomServiceAuthenticationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HealthcareDicomServiceAuthenticationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthcareDicomServiceAuthenticationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HealthcareDicomServiceAuthenticationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHealthcareDicomServiceAuthenticationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

