//go:build no_runtime_type_checking

package virtualmachine

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualMachineStorageDataDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineStorageDataDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineStorageDataDiskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineStorageDataDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineStorageDataDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineStorageDataDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualMachineStorageDataDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

