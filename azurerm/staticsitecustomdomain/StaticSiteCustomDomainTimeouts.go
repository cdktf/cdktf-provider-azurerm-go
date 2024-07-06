// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package staticsitecustomdomain


type StaticSiteCustomDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/static_site_custom_domain#create StaticSiteCustomDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/static_site_custom_domain#delete StaticSiteCustomDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/static_site_custom_domain#read StaticSiteCustomDomain#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

