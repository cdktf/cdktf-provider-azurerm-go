//go:build no_runtime_type_checking

package storageshare

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageShareAclList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageShareAclList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageShareAclListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

