// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_virtual_machine_scale_set#name LinuxVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_virtual_machine_scale_set#product LinuxVirtualMachineScaleSet#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_virtual_machine_scale_set#publisher LinuxVirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
}

