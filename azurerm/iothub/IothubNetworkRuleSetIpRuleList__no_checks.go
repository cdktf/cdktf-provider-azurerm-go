//go:build no_runtime_type_checking

package iothub

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IothubNetworkRuleSetIpRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IothubNetworkRuleSetIpRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IothubNetworkRuleSetIpRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IothubNetworkRuleSetIpRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IothubNetworkRuleSetIpRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IothubNetworkRuleSetIpRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIothubNetworkRuleSetIpRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

