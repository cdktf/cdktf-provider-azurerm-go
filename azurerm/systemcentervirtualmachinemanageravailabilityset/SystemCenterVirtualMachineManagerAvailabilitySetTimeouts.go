// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanageravailabilityset


type SystemCenterVirtualMachineManagerAvailabilitySetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/system_center_virtual_machine_manager_availability_set#create SystemCenterVirtualMachineManagerAvailabilitySet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/system_center_virtual_machine_manager_availability_set#delete SystemCenterVirtualMachineManagerAvailabilitySet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/system_center_virtual_machine_manager_availability_set#read SystemCenterVirtualMachineManagerAvailabilitySet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/system_center_virtual_machine_manager_availability_set#update SystemCenterVirtualMachineManagerAvailabilitySet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

