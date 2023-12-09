// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerappenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerAppEnvironmentWorkloadProfileList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerAppEnvironmentWorkloadProfileList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerAppEnvironmentWorkloadProfileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerAppEnvironmentWorkloadProfileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerAppEnvironmentWorkloadProfileList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerAppEnvironmentWorkloadProfileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerAppEnvironmentWorkloadProfileListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

