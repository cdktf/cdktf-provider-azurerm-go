// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineOsProfileWindowsConfigAdditionalUnattendConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_machine#component VirtualMachine#component}.
	Component *string `field:"required" json:"component" yaml:"component"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_machine#content VirtualMachine#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_machine#pass VirtualMachine#pass}.
	Pass *string `field:"required" json:"pass" yaml:"pass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_machine#setting_name VirtualMachine#setting_name}.
	SettingName *string `field:"required" json:"settingName" yaml:"settingName"`
}

