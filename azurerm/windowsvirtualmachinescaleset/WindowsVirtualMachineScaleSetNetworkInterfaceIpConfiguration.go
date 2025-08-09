// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetNetworkInterfaceIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#application_gateway_backend_address_pool_ids WindowsVirtualMachineScaleSet#application_gateway_backend_address_pool_ids}.
	ApplicationGatewayBackendAddressPoolIds *[]*string `field:"optional" json:"applicationGatewayBackendAddressPoolIds" yaml:"applicationGatewayBackendAddressPoolIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#application_security_group_ids WindowsVirtualMachineScaleSet#application_security_group_ids}.
	ApplicationSecurityGroupIds *[]*string `field:"optional" json:"applicationSecurityGroupIds" yaml:"applicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#load_balancer_backend_address_pool_ids WindowsVirtualMachineScaleSet#load_balancer_backend_address_pool_ids}.
	LoadBalancerBackendAddressPoolIds *[]*string `field:"optional" json:"loadBalancerBackendAddressPoolIds" yaml:"loadBalancerBackendAddressPoolIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#load_balancer_inbound_nat_rules_ids WindowsVirtualMachineScaleSet#load_balancer_inbound_nat_rules_ids}.
	LoadBalancerInboundNatRulesIds *[]*string `field:"optional" json:"loadBalancerInboundNatRulesIds" yaml:"loadBalancerInboundNatRulesIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#primary WindowsVirtualMachineScaleSet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// public_ip_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#public_ip_address WindowsVirtualMachineScaleSet#public_ip_address}
	PublicIpAddress interface{} `field:"optional" json:"publicIpAddress" yaml:"publicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#subnet_id WindowsVirtualMachineScaleSet#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#version WindowsVirtualMachineScaleSet#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

