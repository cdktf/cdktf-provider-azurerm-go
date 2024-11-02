// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetNetworkProfile struct {
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/virtual_machine_scale_set#ip_configuration VirtualMachineScaleSet#ip_configuration}
	IpConfiguration interface{} `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/virtual_machine_scale_set#name VirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/virtual_machine_scale_set#primary VirtualMachineScaleSet#primary}.
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/virtual_machine_scale_set#accelerated_networking VirtualMachineScaleSet#accelerated_networking}.
	AcceleratedNetworking interface{} `field:"optional" json:"acceleratedNetworking" yaml:"acceleratedNetworking"`
	// dns_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/virtual_machine_scale_set#dns_settings VirtualMachineScaleSet#dns_settings}
	DnsSettings *VirtualMachineScaleSetNetworkProfileDnsSettings `field:"optional" json:"dnsSettings" yaml:"dnsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/virtual_machine_scale_set#ip_forwarding VirtualMachineScaleSet#ip_forwarding}.
	IpForwarding interface{} `field:"optional" json:"ipForwarding" yaml:"ipForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/virtual_machine_scale_set#network_security_group_id VirtualMachineScaleSet#network_security_group_id}.
	NetworkSecurityGroupId *string `field:"optional" json:"networkSecurityGroupId" yaml:"networkSecurityGroupId"`
}

