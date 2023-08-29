// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cosmosdbaccount

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CosmosdbAccountGeoLocationList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCosmosdbAccountGeoLocationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

