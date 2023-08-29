// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualmachinescaleset

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualMachineScaleSetExtensionList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualMachineScaleSetExtensionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

