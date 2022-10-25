//go:build no_runtime_type_checking

package containergroup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerGroupInitContainerVolumeList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerGroupInitContainerVolumeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupInitContainerVolumeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupInitContainerVolumeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupInitContainerVolumeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupInitContainerVolumeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerGroupInitContainerVolumeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

