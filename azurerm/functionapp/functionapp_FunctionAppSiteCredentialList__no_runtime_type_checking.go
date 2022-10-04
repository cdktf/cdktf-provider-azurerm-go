//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package functionapp

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionAppSiteCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionAppSiteCredentialList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteCredentialList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSiteCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionAppSiteCredentialListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

