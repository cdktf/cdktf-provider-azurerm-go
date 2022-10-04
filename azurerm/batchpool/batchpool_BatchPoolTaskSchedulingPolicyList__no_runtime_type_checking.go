//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package batchpool

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchPoolTaskSchedulingPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchPoolTaskSchedulingPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchPoolTaskSchedulingPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

