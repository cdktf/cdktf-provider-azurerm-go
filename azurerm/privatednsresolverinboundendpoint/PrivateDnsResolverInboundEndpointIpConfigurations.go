// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatednsresolverinboundendpoint


type PrivateDnsResolverInboundEndpointIpConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/private_dns_resolver_inbound_endpoint#subnet_id PrivateDnsResolverInboundEndpoint#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/private_dns_resolver_inbound_endpoint#private_ip_address PrivateDnsResolverInboundEndpoint#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/private_dns_resolver_inbound_endpoint#private_ip_allocation_method PrivateDnsResolverInboundEndpoint#private_ip_allocation_method}.
	PrivateIpAllocationMethod *string `field:"optional" json:"privateIpAllocationMethod" yaml:"privateIpAllocationMethod"`
}

