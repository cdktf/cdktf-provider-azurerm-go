// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualNetworkIpAddressPoolList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VirtualNetworkIpAddressPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualNetworkIpAddressPoolList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkIpAddressPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkIpAddressPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkIpAddressPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkIpAddressPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualNetworkIpAddressPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

