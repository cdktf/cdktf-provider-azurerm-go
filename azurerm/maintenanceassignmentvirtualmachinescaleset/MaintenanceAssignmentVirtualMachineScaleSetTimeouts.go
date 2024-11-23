// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceassignmentvirtualmachinescaleset


type MaintenanceAssignmentVirtualMachineScaleSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/maintenance_assignment_virtual_machine_scale_set#create MaintenanceAssignmentVirtualMachineScaleSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/maintenance_assignment_virtual_machine_scale_set#delete MaintenanceAssignmentVirtualMachineScaleSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/maintenance_assignment_virtual_machine_scale_set#read MaintenanceAssignmentVirtualMachineScaleSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

