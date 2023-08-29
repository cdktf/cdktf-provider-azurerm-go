// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package trafficmanagerexternalendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TrafficManagerExternalEndpointSubnetList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TrafficManagerExternalEndpointSubnetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TrafficManagerExternalEndpointSubnetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TrafficManagerExternalEndpointSubnetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TrafficManagerExternalEndpointSubnetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TrafficManagerExternalEndpointSubnetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTrafficManagerExternalEndpointSubnetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

