// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/windows_virtual_machine_scale_set#tag WindowsVirtualMachineScaleSet#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/windows_virtual_machine_scale_set#type WindowsVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

