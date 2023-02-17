//go:build no_runtime_type_checking

package batchpool

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchPoolMountCifsMountList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchPoolMountCifsMountList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchPoolMountCifsMountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchPoolMountCifsMountListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

