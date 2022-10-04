//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package storagetable

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageTableAclList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageTableAclList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageTableAclListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

