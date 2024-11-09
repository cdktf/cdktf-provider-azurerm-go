// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cdn_frontdoor_rule#create CdnFrontdoorRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cdn_frontdoor_rule#delete CdnFrontdoorRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cdn_frontdoor_rule#read CdnFrontdoorRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cdn_frontdoor_rule#update CdnFrontdoorRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

