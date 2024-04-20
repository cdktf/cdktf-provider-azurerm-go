// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetNetworkProfileIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_scale_set#name VirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_scale_set#primary VirtualMachineScaleSet#primary}.
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_scale_set#subnet_id VirtualMachineScaleSet#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_scale_set#application_gateway_backend_address_pool_ids VirtualMachineScaleSet#application_gateway_backend_address_pool_ids}.
	ApplicationGatewayBackendAddressPoolIds *[]*string `field:"optional" json:"applicationGatewayBackendAddressPoolIds" yaml:"applicationGatewayBackendAddressPoolIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_scale_set#application_security_group_ids VirtualMachineScaleSet#application_security_group_ids}.
	ApplicationSecurityGroupIds *[]*string `field:"optional" json:"applicationSecurityGroupIds" yaml:"applicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_scale_set#load_balancer_backend_address_pool_ids VirtualMachineScaleSet#load_balancer_backend_address_pool_ids}.
	LoadBalancerBackendAddressPoolIds *[]*string `field:"optional" json:"loadBalancerBackendAddressPoolIds" yaml:"loadBalancerBackendAddressPoolIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_scale_set#load_balancer_inbound_nat_rules_ids VirtualMachineScaleSet#load_balancer_inbound_nat_rules_ids}.
	LoadBalancerInboundNatRulesIds *[]*string `field:"optional" json:"loadBalancerInboundNatRulesIds" yaml:"loadBalancerInboundNatRulesIds"`
	// public_ip_address_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_scale_set#public_ip_address_configuration VirtualMachineScaleSet#public_ip_address_configuration}
	PublicIpAddressConfiguration *VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration `field:"optional" json:"publicIpAddressConfiguration" yaml:"publicIpAddressConfiguration"`
}

