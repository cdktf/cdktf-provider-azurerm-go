// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinerestorepoint


type VirtualMachineRestorePointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/virtual_machine_restore_point#create VirtualMachineRestorePoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/virtual_machine_restore_point#delete VirtualMachineRestorePoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/virtual_machine_restore_point#read VirtualMachineRestorePoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

