// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerappjob

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerAppJobRegistriesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerAppJobRegistriesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerAppJobRegistriesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerAppJobRegistriesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerAppJobRegistriesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerAppJobRegistriesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerAppJobRegistriesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerAppJobRegistriesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

