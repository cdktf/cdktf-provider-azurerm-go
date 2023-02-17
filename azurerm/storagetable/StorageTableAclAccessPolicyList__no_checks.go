//go:build no_runtime_type_checking

package storagetable

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageTableAclAccessPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageTableAclAccessPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclAccessPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclAccessPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclAccessPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclAccessPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageTableAclAccessPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

