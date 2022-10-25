//go:build no_runtime_type_checking

package appconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppConfigurationPrimaryWriteKeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

