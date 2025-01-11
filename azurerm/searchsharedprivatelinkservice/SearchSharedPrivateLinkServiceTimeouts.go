// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchsharedprivatelinkservice


type SearchSharedPrivateLinkServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/search_shared_private_link_service#create SearchSharedPrivateLinkService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/search_shared_private_link_service#delete SearchSharedPrivateLinkService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/search_shared_private_link_service#read SearchSharedPrivateLinkService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/search_shared_private_link_service#update SearchSharedPrivateLinkService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

