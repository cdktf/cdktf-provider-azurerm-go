// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containergroup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerGroupExposedPortList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerGroupExposedPortList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupExposedPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupExposedPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupExposedPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupExposedPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerGroupExposedPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

