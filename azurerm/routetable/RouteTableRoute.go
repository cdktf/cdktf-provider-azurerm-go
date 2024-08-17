// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package routetable


type RouteTableRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/route_table#address_prefix RouteTable#address_prefix}.
	AddressPrefix *string `field:"optional" json:"addressPrefix" yaml:"addressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/route_table#name RouteTable#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/route_table#next_hop_in_ip_address RouteTable#next_hop_in_ip_address}.
	NextHopInIpAddress *string `field:"optional" json:"nextHopInIpAddress" yaml:"nextHopInIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/route_table#next_hop_type RouteTable#next_hop_type}.
	NextHopType *string `field:"optional" json:"nextHopType" yaml:"nextHopType"`
}

