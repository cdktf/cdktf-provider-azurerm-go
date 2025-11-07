// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerconnectedregistry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerConnectedRegistryNotificationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerConnectedRegistryNotificationList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerConnectedRegistryNotificationList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerConnectedRegistryNotificationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerConnectedRegistryNotificationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerConnectedRegistryNotificationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerConnectedRegistryNotificationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerConnectedRegistryNotificationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

