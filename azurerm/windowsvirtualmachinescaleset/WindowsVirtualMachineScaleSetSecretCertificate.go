// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetSecretCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/windows_virtual_machine_scale_set#store WindowsVirtualMachineScaleSet#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/windows_virtual_machine_scale_set#url WindowsVirtualMachineScaleSet#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

