// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#name OrchestratedVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#application_gateway_backend_address_pool_ids OrchestratedVirtualMachineScaleSet#application_gateway_backend_address_pool_ids}.
	ApplicationGatewayBackendAddressPoolIds *[]*string `field:"optional" json:"applicationGatewayBackendAddressPoolIds" yaml:"applicationGatewayBackendAddressPoolIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#application_security_group_ids OrchestratedVirtualMachineScaleSet#application_security_group_ids}.
	ApplicationSecurityGroupIds *[]*string `field:"optional" json:"applicationSecurityGroupIds" yaml:"applicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#load_balancer_backend_address_pool_ids OrchestratedVirtualMachineScaleSet#load_balancer_backend_address_pool_ids}.
	LoadBalancerBackendAddressPoolIds *[]*string `field:"optional" json:"loadBalancerBackendAddressPoolIds" yaml:"loadBalancerBackendAddressPoolIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#primary OrchestratedVirtualMachineScaleSet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// public_ip_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#public_ip_address OrchestratedVirtualMachineScaleSet#public_ip_address}
	PublicIpAddress interface{} `field:"optional" json:"publicIpAddress" yaml:"publicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#subnet_id OrchestratedVirtualMachineScaleSet#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#version OrchestratedVirtualMachineScaleSet#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

