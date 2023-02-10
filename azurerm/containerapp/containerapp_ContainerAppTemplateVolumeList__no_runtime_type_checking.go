//go:build no_runtime_type_checking

package containerapp

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerAppTemplateVolumeList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerAppTemplateVolumeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateVolumeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateVolumeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateVolumeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateVolumeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerAppTemplateVolumeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

