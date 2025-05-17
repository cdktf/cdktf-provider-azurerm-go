// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualhubconnection


type VirtualHubConnectionRoutingStaticVnetRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/virtual_hub_connection#address_prefixes VirtualHubConnection#address_prefixes}.
	AddressPrefixes *[]*string `field:"optional" json:"addressPrefixes" yaml:"addressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/virtual_hub_connection#name VirtualHubConnection#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/virtual_hub_connection#next_hop_ip_address VirtualHubConnection#next_hop_ip_address}.
	NextHopIpAddress *string `field:"optional" json:"nextHopIpAddress" yaml:"nextHopIpAddress"`
}

