// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkconnectionmonitor

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkConnectionMonitorTestConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

