// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualmachine

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualMachineOsProfileSecretsList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineOsProfileSecretsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineOsProfileSecretsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineOsProfileSecretsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineOsProfileSecretsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineOsProfileSecretsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualMachineOsProfileSecretsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

