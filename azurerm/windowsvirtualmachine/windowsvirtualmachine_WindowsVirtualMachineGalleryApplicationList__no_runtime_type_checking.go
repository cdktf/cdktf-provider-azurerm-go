//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package windowsvirtualmachine

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWindowsVirtualMachineGalleryApplicationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

