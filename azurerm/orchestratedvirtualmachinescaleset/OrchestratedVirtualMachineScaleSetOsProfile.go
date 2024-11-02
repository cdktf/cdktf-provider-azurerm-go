// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/orchestrated_virtual_machine_scale_set#custom_data OrchestratedVirtualMachineScaleSet#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// linux_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/orchestrated_virtual_machine_scale_set#linux_configuration OrchestratedVirtualMachineScaleSet#linux_configuration}
	LinuxConfiguration *OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration `field:"optional" json:"linuxConfiguration" yaml:"linuxConfiguration"`
	// windows_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/orchestrated_virtual_machine_scale_set#windows_configuration OrchestratedVirtualMachineScaleSet#windows_configuration}
	WindowsConfiguration *OrchestratedVirtualMachineScaleSetOsProfileWindowsConfiguration `field:"optional" json:"windowsConfiguration" yaml:"windowsConfiguration"`
}

