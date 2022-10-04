//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package frontdoor

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FrontdoorExplicitResourceOrderList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FrontdoorExplicitResourceOrderList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FrontdoorExplicitResourceOrderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFrontdoorExplicitResourceOrderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

