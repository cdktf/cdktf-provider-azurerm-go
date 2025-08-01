// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyvirtualmachineconfigurationassignment


type PolicyVirtualMachineConfigurationAssignmentConfigurationParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/policy_virtual_machine_configuration_assignment#name PolicyVirtualMachineConfigurationAssignment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/policy_virtual_machine_configuration_assignment#value PolicyVirtualMachineConfigurationAssignment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

