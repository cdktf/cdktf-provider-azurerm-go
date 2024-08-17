// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicyrulecollectiongroup


type FirewallPolicyRuleCollectionGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_policy_rule_collection_group#create FirewallPolicyRuleCollectionGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_policy_rule_collection_group#delete FirewallPolicyRuleCollectionGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_policy_rule_collection_group#read FirewallPolicyRuleCollectionGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_policy_rule_collection_group#update FirewallPolicyRuleCollectionGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

