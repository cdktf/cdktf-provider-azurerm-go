//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package windowsfunctionapp

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsFunctionAppSiteCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WindowsFunctionAppSiteCredentialList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppSiteCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppSiteCredentialList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WindowsFunctionAppSiteCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWindowsFunctionAppSiteCredentialListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

