//go:build no_runtime_type_checking

package vmwareprivatecloud

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VmwarePrivateCloudCircuitList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VmwarePrivateCloudCircuitList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VmwarePrivateCloudCircuitList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VmwarePrivateCloudCircuitList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VmwarePrivateCloudCircuitList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVmwarePrivateCloudCircuitListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

