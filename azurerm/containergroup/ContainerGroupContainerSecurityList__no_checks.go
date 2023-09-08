// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containergroup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerGroupContainerSecurityList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerGroupContainerSecurityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupContainerSecurityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupContainerSecurityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupContainerSecurityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupContainerSecurityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerGroupContainerSecurityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

