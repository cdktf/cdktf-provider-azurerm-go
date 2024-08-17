// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallnetworkrulecollection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallNetworkRuleCollectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_network_rule_collection#action FirewallNetworkRuleCollection#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_network_rule_collection#azure_firewall_name FirewallNetworkRuleCollection#azure_firewall_name}.
	AzureFirewallName *string `field:"required" json:"azureFirewallName" yaml:"azureFirewallName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_network_rule_collection#name FirewallNetworkRuleCollection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_network_rule_collection#priority FirewallNetworkRuleCollection#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_network_rule_collection#resource_group_name FirewallNetworkRuleCollection#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_network_rule_collection#rule FirewallNetworkRuleCollection#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_network_rule_collection#id FirewallNetworkRuleCollection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/firewall_network_rule_collection#timeouts FirewallNetworkRuleCollection#timeouts}
	Timeouts *FirewallNetworkRuleCollectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

