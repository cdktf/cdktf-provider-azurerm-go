// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdnendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdnEndpointDeliveryRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCdnEndpointDeliveryRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

