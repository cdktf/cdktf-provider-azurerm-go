//go:build no_runtime_type_checking

package windowswebapp

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsWebAppConnectionStringList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WindowsWebAppConnectionStringList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WindowsWebAppConnectionStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWindowsWebAppConnectionStringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

