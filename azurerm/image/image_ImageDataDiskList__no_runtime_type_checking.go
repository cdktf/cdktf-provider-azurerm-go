//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package image

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImageDataDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_ImageDataDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ImageDataDiskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ImageDataDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ImageDataDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ImageDataDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewImageDataDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

