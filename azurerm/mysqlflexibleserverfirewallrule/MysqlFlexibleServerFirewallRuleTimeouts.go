// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mysqlflexibleserverfirewallrule


type MysqlFlexibleServerFirewallRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mysql_flexible_server_firewall_rule#create MysqlFlexibleServerFirewallRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mysql_flexible_server_firewall_rule#delete MysqlFlexibleServerFirewallRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mysql_flexible_server_firewall_rule#read MysqlFlexibleServerFirewallRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mysql_flexible_server_firewall_rule#update MysqlFlexibleServerFirewallRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

