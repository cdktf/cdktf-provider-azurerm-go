//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataazurermkeyvault

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermKeyVaultAccessPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermKeyVaultAccessPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermKeyVaultAccessPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermKeyVaultAccessPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermKeyVaultAccessPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermKeyVaultAccessPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

