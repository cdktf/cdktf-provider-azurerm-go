// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package windowsvirtualmachine

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsVirtualMachineSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WindowsVirtualMachineSecretList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineSecretList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WindowsVirtualMachineSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWindowsVirtualMachineSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

