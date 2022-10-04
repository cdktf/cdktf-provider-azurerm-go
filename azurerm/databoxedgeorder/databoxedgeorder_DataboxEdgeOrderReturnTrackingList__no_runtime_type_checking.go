//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package databoxedgeorder

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataboxEdgeOrderReturnTrackingList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataboxEdgeOrderReturnTrackingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderReturnTrackingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderReturnTrackingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderReturnTrackingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataboxEdgeOrderReturnTrackingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

