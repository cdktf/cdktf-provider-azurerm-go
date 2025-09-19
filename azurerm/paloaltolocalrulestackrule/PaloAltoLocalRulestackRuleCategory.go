// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackrule


type PaloAltoLocalRulestackRuleCategory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/palo_alto_local_rulestack_rule#custom_urls PaloAltoLocalRulestackRule#custom_urls}.
	CustomUrls *[]*string `field:"required" json:"customUrls" yaml:"customUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/palo_alto_local_rulestack_rule#feeds PaloAltoLocalRulestackRule#feeds}.
	Feeds *[]*string `field:"optional" json:"feeds" yaml:"feeds"`
}

