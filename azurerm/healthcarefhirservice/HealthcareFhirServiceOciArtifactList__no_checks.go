// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package healthcarefhirservice

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthcareFhirServiceOciArtifactList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HealthcareFhirServiceOciArtifactList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HealthcareFhirServiceOciArtifactList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HealthcareFhirServiceOciArtifactList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthcareFhirServiceOciArtifactList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HealthcareFhirServiceOciArtifactList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHealthcareFhirServiceOciArtifactListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

