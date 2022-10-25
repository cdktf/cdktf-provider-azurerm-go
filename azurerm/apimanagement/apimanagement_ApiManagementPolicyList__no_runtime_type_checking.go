//go:build no_runtime_type_checking

package apimanagement

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiManagementPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiManagementPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiManagementPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiManagementPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiManagementPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiManagementPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiManagementPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

