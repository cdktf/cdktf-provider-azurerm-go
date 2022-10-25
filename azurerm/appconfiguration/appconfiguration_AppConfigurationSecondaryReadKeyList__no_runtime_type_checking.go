//go:build no_runtime_type_checking

package appconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppConfigurationSecondaryReadKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppConfigurationSecondaryReadKeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationSecondaryReadKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationSecondaryReadKeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationSecondaryReadKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppConfigurationSecondaryReadKeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

