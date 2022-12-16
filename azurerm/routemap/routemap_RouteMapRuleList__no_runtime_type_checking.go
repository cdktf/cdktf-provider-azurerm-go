//go:build no_runtime_type_checking

package routemap

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteMapRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RouteMapRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRouteMapRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

