// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineruncommand


type VirtualMachineRunCommandParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/virtual_machine_run_command#name VirtualMachineRunCommand#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/virtual_machine_run_command#value VirtualMachineRunCommand#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

