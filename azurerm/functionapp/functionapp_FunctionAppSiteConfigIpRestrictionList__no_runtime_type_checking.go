//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package functionapp

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionAppSiteConfigIpRestrictionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

