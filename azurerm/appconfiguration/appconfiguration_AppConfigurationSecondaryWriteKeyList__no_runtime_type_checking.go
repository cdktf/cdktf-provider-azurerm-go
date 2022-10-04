//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package appconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppConfigurationSecondaryWriteKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppConfigurationSecondaryWriteKeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationSecondaryWriteKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationSecondaryWriteKeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationSecondaryWriteKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppConfigurationSecondaryWriteKeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

