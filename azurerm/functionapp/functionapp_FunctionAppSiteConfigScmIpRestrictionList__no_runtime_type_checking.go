//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package functionapp

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionAppSiteConfigScmIpRestrictionList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionAppSiteConfigScmIpRestrictionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteConfigScmIpRestrictionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteConfigScmIpRestrictionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteConfigScmIpRestrictionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteConfigScmIpRestrictionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionAppSiteConfigScmIpRestrictionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

