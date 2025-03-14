// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetAutomaticInstanceRepair struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/linux_virtual_machine_scale_set#enabled LinuxVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/linux_virtual_machine_scale_set#action LinuxVirtualMachineScaleSet#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/linux_virtual_machine_scale_set#grace_period LinuxVirtualMachineScaleSet#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}

