//go:build no_runtime_type_checking

package linuxfunctionapp

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxFunctionAppConnectionStringList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LinuxFunctionAppConnectionStringList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppConnectionStringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppConnectionStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppConnectionStringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppConnectionStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLinuxFunctionAppConnectionStringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

