// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetTerminateNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_virtual_machine_scale_set#enabled LinuxVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_virtual_machine_scale_set#timeout LinuxVirtualMachineScaleSet#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

