// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package linuxvirtualmachine

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxVirtualMachineSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LinuxVirtualMachineSecretList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LinuxVirtualMachineSecretList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LinuxVirtualMachineSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LinuxVirtualMachineSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LinuxVirtualMachineSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLinuxVirtualMachineSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

