// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lb


type LbFrontendIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#name Lb#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#gateway_load_balancer_frontend_ip_configuration_id Lb#gateway_load_balancer_frontend_ip_configuration_id}.
	GatewayLoadBalancerFrontendIpConfigurationId *string `field:"optional" json:"gatewayLoadBalancerFrontendIpConfigurationId" yaml:"gatewayLoadBalancerFrontendIpConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#private_ip_address Lb#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#private_ip_address_allocation Lb#private_ip_address_allocation}.
	PrivateIpAddressAllocation *string `field:"optional" json:"privateIpAddressAllocation" yaml:"privateIpAddressAllocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#private_ip_address_version Lb#private_ip_address_version}.
	PrivateIpAddressVersion *string `field:"optional" json:"privateIpAddressVersion" yaml:"privateIpAddressVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#public_ip_address_id Lb#public_ip_address_id}.
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#public_ip_prefix_id Lb#public_ip_prefix_id}.
	PublicIpPrefixId *string `field:"optional" json:"publicIpPrefixId" yaml:"publicIpPrefixId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#subnet_id Lb#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/lb#zones Lb#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

