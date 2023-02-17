//go:build no_runtime_type_checking

package appserviceenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppServiceEnvironmentClusterSettingList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppServiceEnvironmentClusterSettingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppServiceEnvironmentClusterSettingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppServiceEnvironmentClusterSettingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppServiceEnvironmentClusterSettingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppServiceEnvironmentClusterSettingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppServiceEnvironmentClusterSettingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

