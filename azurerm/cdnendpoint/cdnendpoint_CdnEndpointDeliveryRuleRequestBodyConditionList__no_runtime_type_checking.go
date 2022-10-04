//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cdnendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleRequestBodyConditionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCdnEndpointDeliveryRuleRequestBodyConditionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

