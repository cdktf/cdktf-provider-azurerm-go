// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mongoclusterfirewallrule


type MongoClusterFirewallRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/mongo_cluster_firewall_rule#create MongoClusterFirewallRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/mongo_cluster_firewall_rule#delete MongoClusterFirewallRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/mongo_cluster_firewall_rule#read MongoClusterFirewallRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/mongo_cluster_firewall_rule#update MongoClusterFirewallRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

