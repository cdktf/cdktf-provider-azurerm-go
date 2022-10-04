//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package linuxfunctionapp

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxFunctionAppSiteCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LinuxFunctionAppSiteCredentialList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppSiteCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppSiteCredentialList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppSiteCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLinuxFunctionAppSiteCredentialListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

