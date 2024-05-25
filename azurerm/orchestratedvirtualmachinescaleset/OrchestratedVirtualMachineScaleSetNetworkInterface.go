// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetNetworkInterface struct {
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/orchestrated_virtual_machine_scale_set#ip_configuration OrchestratedVirtualMachineScaleSet#ip_configuration}
	IpConfiguration interface{} `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/orchestrated_virtual_machine_scale_set#name OrchestratedVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/orchestrated_virtual_machine_scale_set#dns_servers OrchestratedVirtualMachineScaleSet#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/orchestrated_virtual_machine_scale_set#enable_accelerated_networking OrchestratedVirtualMachineScaleSet#enable_accelerated_networking}.
	EnableAcceleratedNetworking interface{} `field:"optional" json:"enableAcceleratedNetworking" yaml:"enableAcceleratedNetworking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/orchestrated_virtual_machine_scale_set#enable_ip_forwarding OrchestratedVirtualMachineScaleSet#enable_ip_forwarding}.
	EnableIpForwarding interface{} `field:"optional" json:"enableIpForwarding" yaml:"enableIpForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/orchestrated_virtual_machine_scale_set#network_security_group_id OrchestratedVirtualMachineScaleSet#network_security_group_id}.
	NetworkSecurityGroupId *string `field:"optional" json:"networkSecurityGroupId" yaml:"networkSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/orchestrated_virtual_machine_scale_set#primary OrchestratedVirtualMachineScaleSet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

