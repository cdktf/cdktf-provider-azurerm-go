// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicyrulecollectiongroup


type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocols struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/firewall_policy_rule_collection_group#port FirewallPolicyRuleCollectionGroup#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/firewall_policy_rule_collection_group#type FirewallPolicyRuleCollectionGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

