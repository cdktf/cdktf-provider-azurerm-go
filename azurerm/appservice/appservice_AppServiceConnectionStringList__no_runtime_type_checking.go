//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package appservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppServiceConnectionStringList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppServiceConnectionStringList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppServiceConnectionStringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppServiceConnectionStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppServiceConnectionStringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppServiceConnectionStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppServiceConnectionStringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

