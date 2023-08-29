// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkconnectionmonitor

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkConnectionMonitorEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkConnectionMonitorEndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

