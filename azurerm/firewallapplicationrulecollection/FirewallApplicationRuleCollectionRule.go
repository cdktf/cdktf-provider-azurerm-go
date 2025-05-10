// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallapplicationrulecollection


type FirewallApplicationRuleCollectionRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/firewall_application_rule_collection#name FirewallApplicationRuleCollection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/firewall_application_rule_collection#description FirewallApplicationRuleCollection#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/firewall_application_rule_collection#fqdn_tags FirewallApplicationRuleCollection#fqdn_tags}.
	FqdnTags *[]*string `field:"optional" json:"fqdnTags" yaml:"fqdnTags"`
	// protocol block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/firewall_application_rule_collection#protocol FirewallApplicationRuleCollection#protocol}
	Protocol interface{} `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/firewall_application_rule_collection#source_addresses FirewallApplicationRuleCollection#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/firewall_application_rule_collection#source_ip_groups FirewallApplicationRuleCollection#source_ip_groups}.
	SourceIpGroups *[]*string `field:"optional" json:"sourceIpGroups" yaml:"sourceIpGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/firewall_application_rule_collection#target_fqdns FirewallApplicationRuleCollection#target_fqdns}.
	TargetFqdns *[]*string `field:"optional" json:"targetFqdns" yaml:"targetFqdns"`
}

