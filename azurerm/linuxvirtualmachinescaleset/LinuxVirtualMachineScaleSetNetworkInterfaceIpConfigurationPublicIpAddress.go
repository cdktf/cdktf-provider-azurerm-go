// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/linux_virtual_machine_scale_set#name LinuxVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/linux_virtual_machine_scale_set#domain_name_label LinuxVirtualMachineScaleSet#domain_name_label}.
	DomainNameLabel *string `field:"optional" json:"domainNameLabel" yaml:"domainNameLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/linux_virtual_machine_scale_set#idle_timeout_in_minutes LinuxVirtualMachineScaleSet#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// ip_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/linux_virtual_machine_scale_set#ip_tag LinuxVirtualMachineScaleSet#ip_tag}
	IpTag interface{} `field:"optional" json:"ipTag" yaml:"ipTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/linux_virtual_machine_scale_set#public_ip_prefix_id LinuxVirtualMachineScaleSet#public_ip_prefix_id}.
	PublicIpPrefixId *string `field:"optional" json:"publicIpPrefixId" yaml:"publicIpPrefixId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/linux_virtual_machine_scale_set#version LinuxVirtualMachineScaleSet#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

