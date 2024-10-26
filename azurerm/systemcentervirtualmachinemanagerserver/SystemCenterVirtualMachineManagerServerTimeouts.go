// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagerserver


type SystemCenterVirtualMachineManagerServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/system_center_virtual_machine_manager_server#create SystemCenterVirtualMachineManagerServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/system_center_virtual_machine_manager_server#delete SystemCenterVirtualMachineManagerServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/system_center_virtual_machine_manager_server#read SystemCenterVirtualMachineManagerServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/system_center_virtual_machine_manager_server#update SystemCenterVirtualMachineManagerServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

