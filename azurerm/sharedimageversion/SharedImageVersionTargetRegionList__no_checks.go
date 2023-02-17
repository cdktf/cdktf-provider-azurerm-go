//go:build no_runtime_type_checking

package sharedimageversion

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SharedImageVersionTargetRegionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SharedImageVersionTargetRegionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SharedImageVersionTargetRegionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImageVersionTargetRegionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImageVersionTargetRegionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SharedImageVersionTargetRegionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSharedImageVersionTargetRegionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

