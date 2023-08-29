// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdnendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRulePostArgConditionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCdnEndpointDeliveryRulePostArgConditionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

