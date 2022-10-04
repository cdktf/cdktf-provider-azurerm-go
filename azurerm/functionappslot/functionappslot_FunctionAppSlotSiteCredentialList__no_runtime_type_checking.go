//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package functionappslot

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionAppSlotSiteCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionAppSlotSiteCredentialList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSlotSiteCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSlotSiteCredentialList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionAppSlotSiteCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionAppSlotSiteCredentialListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

