//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cdnfrontdoorrule

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdnFrontdoorRuleActionsRequestHeaderActionList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRequestHeaderActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRequestHeaderActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRequestHeaderActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRequestHeaderActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRequestHeaderActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCdnFrontdoorRuleActionsRequestHeaderActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
