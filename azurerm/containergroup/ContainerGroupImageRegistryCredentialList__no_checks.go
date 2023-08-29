// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containergroup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerGroupImageRegistryCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerGroupImageRegistryCredentialList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupImageRegistryCredentialList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupImageRegistryCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupImageRegistryCredentialList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerGroupImageRegistryCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerGroupImageRegistryCredentialListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

