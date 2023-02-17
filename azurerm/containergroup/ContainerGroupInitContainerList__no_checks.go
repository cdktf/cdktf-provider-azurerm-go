//go:build no_runtime_type_checking

package containergroup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerGroupInitContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerGroupInitContainerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupInitContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupInitContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupInitContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupInitContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerGroupInitContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

