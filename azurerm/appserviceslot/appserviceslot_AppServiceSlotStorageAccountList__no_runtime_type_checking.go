//go:build no_runtime_type_checking

package appserviceslot

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppServiceSlotStorageAccountList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppServiceSlotStorageAccountList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppServiceSlotStorageAccountListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

