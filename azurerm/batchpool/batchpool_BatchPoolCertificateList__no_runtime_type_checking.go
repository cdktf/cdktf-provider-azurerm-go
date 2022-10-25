//go:build no_runtime_type_checking

package batchpool

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchPoolCertificateList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchPoolCertificateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchPoolCertificateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchPoolCertificateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchPoolCertificateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchPoolCertificateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchPoolCertificateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

