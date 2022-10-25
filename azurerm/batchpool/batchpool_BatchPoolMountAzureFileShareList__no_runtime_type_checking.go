//go:build no_runtime_type_checking

package batchpool

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchPoolMountAzureFileShareList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchPoolMountAzureFileShareList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchPoolMountAzureFileShareList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchPoolMountAzureFileShareListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

