//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package databoxedgeorder

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataboxEdgeOrderStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataboxEdgeOrderStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataboxEdgeOrderStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

