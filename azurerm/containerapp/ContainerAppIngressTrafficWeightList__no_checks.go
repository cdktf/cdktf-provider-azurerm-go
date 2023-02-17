//go:build no_runtime_type_checking

package containerapp

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerAppIngressTrafficWeightList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerAppIngressTrafficWeightList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerAppIngressTrafficWeightList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerAppIngressTrafficWeightList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerAppIngressTrafficWeightList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerAppIngressTrafficWeightList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerAppIngressTrafficWeightListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

