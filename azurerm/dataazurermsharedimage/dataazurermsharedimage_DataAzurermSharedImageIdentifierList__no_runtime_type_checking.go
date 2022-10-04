//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataazurermsharedimage

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermSharedImageIdentifierList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermSharedImageIdentifierList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermSharedImageIdentifierList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermSharedImageIdentifierList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermSharedImageIdentifierList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermSharedImageIdentifierListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

