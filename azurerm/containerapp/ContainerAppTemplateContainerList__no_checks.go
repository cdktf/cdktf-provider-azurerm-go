//go:build no_runtime_type_checking

package containerapp

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerAppTemplateContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerAppTemplateContainerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerAppTemplateContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
