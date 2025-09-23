// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine


type WindowsVirtualMachineTerminationNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/windows_virtual_machine#enabled WindowsVirtualMachine#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/windows_virtual_machine#timeout WindowsVirtualMachine#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

