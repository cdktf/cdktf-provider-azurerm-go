// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_virtual_machine_scale_set#tag LinuxVirtualMachineScaleSet#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_virtual_machine_scale_set#type LinuxVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

