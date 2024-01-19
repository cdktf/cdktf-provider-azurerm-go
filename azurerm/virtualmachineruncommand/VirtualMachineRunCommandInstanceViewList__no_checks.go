// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualmachineruncommand

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualMachineRunCommandInstanceViewList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineRunCommandInstanceViewList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineRunCommandInstanceViewList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineRunCommandInstanceViewList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineRunCommandInstanceViewList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineRunCommandInstanceViewList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualMachineRunCommandInstanceViewListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

