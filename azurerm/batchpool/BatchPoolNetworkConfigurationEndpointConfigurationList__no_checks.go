// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package batchpool

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchPoolNetworkConfigurationEndpointConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

