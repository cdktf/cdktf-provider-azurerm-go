// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualmachineinstance


type SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystem struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#computer_name SystemCenterVirtualMachineManagerVirtualMachineInstance#computer_name}.
	ComputerName *string `field:"required" json:"computerName" yaml:"computerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#admin_password SystemCenterVirtualMachineManagerVirtualMachineInstance#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
}

