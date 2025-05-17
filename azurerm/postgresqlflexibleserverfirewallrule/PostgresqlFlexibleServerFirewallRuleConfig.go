// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlflexibleserverfirewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PostgresqlFlexibleServerFirewallRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/postgresql_flexible_server_firewall_rule#end_ip_address PostgresqlFlexibleServerFirewallRule#end_ip_address}.
	EndIpAddress *string `field:"required" json:"endIpAddress" yaml:"endIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/postgresql_flexible_server_firewall_rule#name PostgresqlFlexibleServerFirewallRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/postgresql_flexible_server_firewall_rule#server_id PostgresqlFlexibleServerFirewallRule#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/postgresql_flexible_server_firewall_rule#start_ip_address PostgresqlFlexibleServerFirewallRule#start_ip_address}.
	StartIpAddress *string `field:"required" json:"startIpAddress" yaml:"startIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/postgresql_flexible_server_firewall_rule#id PostgresqlFlexibleServerFirewallRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/postgresql_flexible_server_firewall_rule#timeouts PostgresqlFlexibleServerFirewallRule#timeouts}
	Timeouts *PostgresqlFlexibleServerFirewallRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

