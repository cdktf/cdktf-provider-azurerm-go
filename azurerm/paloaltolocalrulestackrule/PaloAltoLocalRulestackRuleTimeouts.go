// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackrule


type PaloAltoLocalRulestackRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/palo_alto_local_rulestack_rule#create PaloAltoLocalRulestackRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/palo_alto_local_rulestack_rule#delete PaloAltoLocalRulestackRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/palo_alto_local_rulestack_rule#read PaloAltoLocalRulestackRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/palo_alto_local_rulestack_rule#update PaloAltoLocalRulestackRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

