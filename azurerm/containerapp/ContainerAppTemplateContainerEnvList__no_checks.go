// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerapp

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerAppTemplateContainerEnvList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerAppTemplateContainerEnvList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerEnvList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerEnvList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerEnvList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerEnvList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerAppTemplateContainerEnvListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

