// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine


type WindowsVirtualMachineOsDiskDiffDiskSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_virtual_machine#option WindowsVirtualMachine#option}.
	Option *string `field:"required" json:"option" yaml:"option"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_virtual_machine#placement WindowsVirtualMachine#placement}.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
}

