// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatednsresolvervirtualnetworklink


type PrivateDnsResolverVirtualNetworkLinkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/private_dns_resolver_virtual_network_link#create PrivateDnsResolverVirtualNetworkLink#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/private_dns_resolver_virtual_network_link#delete PrivateDnsResolverVirtualNetworkLink#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/private_dns_resolver_virtual_network_link#read PrivateDnsResolverVirtualNetworkLink#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/private_dns_resolver_virtual_network_link#update PrivateDnsResolverVirtualNetworkLink#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

