// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerregistry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRegistryRetentionPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryRetentionPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryRetentionPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryRetentionPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryRetentionPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryRetentionPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryRetentionPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerRegistryRetentionPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

