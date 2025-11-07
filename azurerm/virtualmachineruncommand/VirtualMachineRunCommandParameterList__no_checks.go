// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualmachineruncommand

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualMachineRunCommandParameterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineRunCommandParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineRunCommandParameterList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineRunCommandParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineRunCommandParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineRunCommandParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineRunCommandParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualMachineRunCommandParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

