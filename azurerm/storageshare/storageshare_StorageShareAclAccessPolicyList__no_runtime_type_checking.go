//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package storageshare

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageShareAclAccessPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageShareAclAccessPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclAccessPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclAccessPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclAccessPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclAccessPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageShareAclAccessPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

