// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetNetworkInterface struct {
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#ip_configuration WindowsVirtualMachineScaleSet#ip_configuration}
	IpConfiguration interface{} `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#auxiliary_mode WindowsVirtualMachineScaleSet#auxiliary_mode}.
	AuxiliaryMode *string `field:"optional" json:"auxiliaryMode" yaml:"auxiliaryMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#auxiliary_sku WindowsVirtualMachineScaleSet#auxiliary_sku}.
	AuxiliarySku *string `field:"optional" json:"auxiliarySku" yaml:"auxiliarySku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#dns_servers WindowsVirtualMachineScaleSet#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#enable_accelerated_networking WindowsVirtualMachineScaleSet#enable_accelerated_networking}.
	EnableAcceleratedNetworking interface{} `field:"optional" json:"enableAcceleratedNetworking" yaml:"enableAcceleratedNetworking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#enable_ip_forwarding WindowsVirtualMachineScaleSet#enable_ip_forwarding}.
	EnableIpForwarding interface{} `field:"optional" json:"enableIpForwarding" yaml:"enableIpForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#network_security_group_id WindowsVirtualMachineScaleSet#network_security_group_id}.
	NetworkSecurityGroupId *string `field:"optional" json:"networkSecurityGroupId" yaml:"networkSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/windows_virtual_machine_scale_set#primary WindowsVirtualMachineScaleSet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

