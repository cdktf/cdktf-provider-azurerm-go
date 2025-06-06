// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vpngateway

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpnGatewayIpConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VpnGatewayIpConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VpnGatewayIpConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VpnGatewayIpConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VpnGatewayIpConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VpnGatewayIpConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVpnGatewayIpConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

