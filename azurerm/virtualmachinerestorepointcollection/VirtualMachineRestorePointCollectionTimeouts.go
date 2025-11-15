// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinerestorepointcollection


type VirtualMachineRestorePointCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/virtual_machine_restore_point_collection#create VirtualMachineRestorePointCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/virtual_machine_restore_point_collection#delete VirtualMachineRestorePointCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/virtual_machine_restore_point_collection#read VirtualMachineRestorePointCollection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/virtual_machine_restore_point_collection#update VirtualMachineRestorePointCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

