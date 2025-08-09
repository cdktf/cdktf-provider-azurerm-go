// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackrule


type PaloAltoLocalRulestackRuleSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/palo_alto_local_rulestack_rule#cidrs PaloAltoLocalRulestackRule#cidrs}.
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/palo_alto_local_rulestack_rule#countries PaloAltoLocalRulestackRule#countries}.
	Countries *[]*string `field:"optional" json:"countries" yaml:"countries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/palo_alto_local_rulestack_rule#feeds PaloAltoLocalRulestackRule#feeds}.
	Feeds *[]*string `field:"optional" json:"feeds" yaml:"feeds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/palo_alto_local_rulestack_rule#local_rulestack_prefix_list_ids PaloAltoLocalRulestackRule#local_rulestack_prefix_list_ids}.
	LocalRulestackPrefixListIds *[]*string `field:"optional" json:"localRulestackPrefixListIds" yaml:"localRulestackPrefixListIds"`
}

