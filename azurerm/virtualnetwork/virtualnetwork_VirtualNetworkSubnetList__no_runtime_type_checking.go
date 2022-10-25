//go:build no_runtime_type_checking

package virtualnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualNetworkSubnetList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualNetworkSubnetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualNetworkSubnetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

