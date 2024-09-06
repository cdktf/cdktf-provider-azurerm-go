// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/virtual_machine_scale_set#domain_name_label VirtualMachineScaleSet#domain_name_label}.
	DomainNameLabel *string `field:"required" json:"domainNameLabel" yaml:"domainNameLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/virtual_machine_scale_set#idle_timeout VirtualMachineScaleSet#idle_timeout}.
	IdleTimeout *float64 `field:"required" json:"idleTimeout" yaml:"idleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/virtual_machine_scale_set#name VirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

