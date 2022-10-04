//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package windowsvirtualmachine

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsVirtualMachineSecretCertificateList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WindowsVirtualMachineSecretCertificateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineSecretCertificateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineSecretCertificateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineSecretCertificateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineSecretCertificateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWindowsVirtualMachineSecretCertificateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

