//go:build no_runtime_type_checking

package batchpool

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchPoolStartTaskContainerRegistryList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchPoolStartTaskContainerRegistryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchPoolStartTaskContainerRegistryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchPoolStartTaskContainerRegistryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchPoolStartTaskContainerRegistryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchPoolStartTaskContainerRegistryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchPoolStartTaskContainerRegistryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
