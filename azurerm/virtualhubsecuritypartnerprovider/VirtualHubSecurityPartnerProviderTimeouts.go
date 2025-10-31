// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualhubsecuritypartnerprovider


type VirtualHubSecurityPartnerProviderTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_hub_security_partner_provider#create VirtualHubSecurityPartnerProvider#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_hub_security_partner_provider#delete VirtualHubSecurityPartnerProvider#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_hub_security_partner_provider#read VirtualHubSecurityPartnerProvider#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_hub_security_partner_provider#update VirtualHubSecurityPartnerProvider#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

