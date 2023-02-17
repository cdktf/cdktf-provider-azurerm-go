//go:build no_runtime_type_checking

package frontdoor

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FrontdoorBackendPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FrontdoorBackendPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFrontdoorBackendPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

