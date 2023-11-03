// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicy


type FirewallPolicyIntrusionDetectionTrafficBypass struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/firewall_policy#name FirewallPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/firewall_policy#protocol FirewallPolicy#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/firewall_policy#description FirewallPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/firewall_policy#destination_addresses FirewallPolicy#destination_addresses}.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/firewall_policy#destination_ip_groups FirewallPolicy#destination_ip_groups}.
	DestinationIpGroups *[]*string `field:"optional" json:"destinationIpGroups" yaml:"destinationIpGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/firewall_policy#destination_ports FirewallPolicy#destination_ports}.
	DestinationPorts *[]*string `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/firewall_policy#source_addresses FirewallPolicy#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/firewall_policy#source_ip_groups FirewallPolicy#source_ip_groups}.
	SourceIpGroups *[]*string `field:"optional" json:"sourceIpGroups" yaml:"sourceIpGroups"`
}

