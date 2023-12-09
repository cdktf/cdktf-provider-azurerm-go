// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorruleset


type CdnFrontdoorRuleSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/cdn_frontdoor_rule_set#create CdnFrontdoorRuleSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/cdn_frontdoor_rule_set#delete CdnFrontdoorRuleSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/cdn_frontdoor_rule_set#read CdnFrontdoorRuleSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

