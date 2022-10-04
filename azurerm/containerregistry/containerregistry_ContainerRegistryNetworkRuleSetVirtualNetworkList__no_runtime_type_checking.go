//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package containerregistry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRegistryNetworkRuleSetVirtualNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryNetworkRuleSetVirtualNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryNetworkRuleSetVirtualNetworkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryNetworkRuleSetVirtualNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryNetworkRuleSetVirtualNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryNetworkRuleSetVirtualNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerRegistryNetworkRuleSetVirtualNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

