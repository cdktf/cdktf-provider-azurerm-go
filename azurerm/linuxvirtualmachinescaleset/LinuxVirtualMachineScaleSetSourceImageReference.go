// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetSourceImageReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/linux_virtual_machine_scale_set#offer LinuxVirtualMachineScaleSet#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/linux_virtual_machine_scale_set#publisher LinuxVirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/linux_virtual_machine_scale_set#sku LinuxVirtualMachineScaleSet#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/linux_virtual_machine_scale_set#version LinuxVirtualMachineScaleSet#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

