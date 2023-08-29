// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualhub

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualHubRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualHubRouteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualHubRouteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

