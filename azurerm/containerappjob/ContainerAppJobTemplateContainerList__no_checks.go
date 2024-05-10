// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerappjob

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerAppJobTemplateContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerAppJobTemplateContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerAppJobTemplateContainerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerAppJobTemplateContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerAppJobTemplateContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerAppJobTemplateContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerAppJobTemplateContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerAppJobTemplateContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

