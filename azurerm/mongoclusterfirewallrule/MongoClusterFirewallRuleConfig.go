// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mongoclusterfirewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongoClusterFirewallRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/mongo_cluster_firewall_rule#end_ip_address MongoClusterFirewallRule#end_ip_address}.
	EndIpAddress *string `field:"required" json:"endIpAddress" yaml:"endIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/mongo_cluster_firewall_rule#mongo_cluster_id MongoClusterFirewallRule#mongo_cluster_id}.
	MongoClusterId *string `field:"required" json:"mongoClusterId" yaml:"mongoClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/mongo_cluster_firewall_rule#name MongoClusterFirewallRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/mongo_cluster_firewall_rule#start_ip_address MongoClusterFirewallRule#start_ip_address}.
	StartIpAddress *string `field:"required" json:"startIpAddress" yaml:"startIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/mongo_cluster_firewall_rule#id MongoClusterFirewallRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/mongo_cluster_firewall_rule#timeouts MongoClusterFirewallRule#timeouts}
	Timeouts *MongoClusterFirewallRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

