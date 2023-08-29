// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iothub

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IothubEnrichmentList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IothubEnrichmentList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IothubEnrichmentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IothubEnrichmentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IothubEnrichmentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IothubEnrichmentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIothubEnrichmentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

