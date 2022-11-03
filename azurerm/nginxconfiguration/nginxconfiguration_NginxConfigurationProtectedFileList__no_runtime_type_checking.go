//go:build no_runtime_type_checking

package nginxconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NginxConfigurationProtectedFileList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NginxConfigurationProtectedFileList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationProtectedFileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationProtectedFileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationProtectedFileList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationProtectedFileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNginxConfigurationProtectedFileListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

