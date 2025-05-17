// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallapplicationrulecollection


type FirewallApplicationRuleCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/firewall_application_rule_collection#create FirewallApplicationRuleCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/firewall_application_rule_collection#delete FirewallApplicationRuleCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/firewall_application_rule_collection#read FirewallApplicationRuleCollection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/firewall_application_rule_collection#update FirewallApplicationRuleCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

