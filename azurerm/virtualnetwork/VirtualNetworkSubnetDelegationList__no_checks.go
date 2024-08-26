// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualNetworkSubnetDelegationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VirtualNetworkSubnetDelegationList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualNetworkSubnetDelegationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetDelegationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetDelegationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetDelegationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetDelegationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualNetworkSubnetDelegationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

