//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataazurermbatchpool

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermBatchPoolExtensionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermBatchPoolExtensionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermBatchPoolExtensionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermBatchPoolExtensionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermBatchPoolExtensionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermBatchPoolExtensionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
