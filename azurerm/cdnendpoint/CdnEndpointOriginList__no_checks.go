// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdnendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdnEndpointOriginList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CdnEndpointOriginList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CdnEndpointOriginList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointOriginList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointOriginList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointOriginList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointOriginList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCdnEndpointOriginListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

