//go:build no_runtime_type_checking

package appservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppServiceSiteConfigIpRestrictionHeadersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppServiceSiteConfigIpRestrictionHeadersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteConfigIpRestrictionHeadersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteConfigIpRestrictionHeadersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteConfigIpRestrictionHeadersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteConfigIpRestrictionHeadersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppServiceSiteConfigIpRestrictionHeadersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

