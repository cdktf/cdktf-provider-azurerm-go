//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package linuxvirtualmachine

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLinuxVirtualMachineGalleryApplicationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

