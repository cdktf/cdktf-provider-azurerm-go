//go:build no_runtime_type_checking

package frontdoor

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolLoadBalancingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFrontdoorBackendPoolLoadBalancingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

