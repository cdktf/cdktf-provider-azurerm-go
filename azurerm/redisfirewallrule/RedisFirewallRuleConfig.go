// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redisfirewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedisFirewallRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/redis_firewall_rule#end_ip RedisFirewallRule#end_ip}.
	EndIp *string `field:"required" json:"endIp" yaml:"endIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/redis_firewall_rule#name RedisFirewallRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/redis_firewall_rule#redis_cache_name RedisFirewallRule#redis_cache_name}.
	RedisCacheName *string `field:"required" json:"redisCacheName" yaml:"redisCacheName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/redis_firewall_rule#resource_group_name RedisFirewallRule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/redis_firewall_rule#start_ip RedisFirewallRule#start_ip}.
	StartIp *string `field:"required" json:"startIp" yaml:"startIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/redis_firewall_rule#id RedisFirewallRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/redis_firewall_rule#timeouts RedisFirewallRule#timeouts}
	Timeouts *RedisFirewallRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

