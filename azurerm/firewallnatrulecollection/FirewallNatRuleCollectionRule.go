// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallnatrulecollection


type FirewallNatRuleCollectionRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#destination_addresses FirewallNatRuleCollection#destination_addresses}.
	DestinationAddresses *[]*string `field:"required" json:"destinationAddresses" yaml:"destinationAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#destination_ports FirewallNatRuleCollection#destination_ports}.
	DestinationPorts *[]*string `field:"required" json:"destinationPorts" yaml:"destinationPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#name FirewallNatRuleCollection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#protocols FirewallNatRuleCollection#protocols}.
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#translated_address FirewallNatRuleCollection#translated_address}.
	TranslatedAddress *string `field:"required" json:"translatedAddress" yaml:"translatedAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#translated_port FirewallNatRuleCollection#translated_port}.
	TranslatedPort *string `field:"required" json:"translatedPort" yaml:"translatedPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#description FirewallNatRuleCollection#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#source_addresses FirewallNatRuleCollection#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/firewall_nat_rule_collection#source_ip_groups FirewallNatRuleCollection#source_ip_groups}.
	SourceIpGroups *[]*string `field:"optional" json:"sourceIpGroups" yaml:"sourceIpGroups"`
}

