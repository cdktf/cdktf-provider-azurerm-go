// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestlinuxvirtualmachine


type DevTestLinuxVirtualMachineGalleryImageReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_test_linux_virtual_machine#offer DevTestLinuxVirtualMachine#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_test_linux_virtual_machine#publisher DevTestLinuxVirtualMachine#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_test_linux_virtual_machine#sku DevTestLinuxVirtualMachine#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_test_linux_virtual_machine#version DevTestLinuxVirtualMachine#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

