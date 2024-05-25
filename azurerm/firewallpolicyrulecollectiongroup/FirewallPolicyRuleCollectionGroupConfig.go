// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicyrulecollectiongroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallPolicyRuleCollectionGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/firewall_policy_rule_collection_group#firewall_policy_id FirewallPolicyRuleCollectionGroup#firewall_policy_id}.
	FirewallPolicyId *string `field:"required" json:"firewallPolicyId" yaml:"firewallPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/firewall_policy_rule_collection_group#priority FirewallPolicyRuleCollectionGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// application_rule_collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/firewall_policy_rule_collection_group#application_rule_collection FirewallPolicyRuleCollectionGroup#application_rule_collection}
	ApplicationRuleCollection interface{} `field:"optional" json:"applicationRuleCollection" yaml:"applicationRuleCollection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/firewall_policy_rule_collection_group#id FirewallPolicyRuleCollectionGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// nat_rule_collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/firewall_policy_rule_collection_group#nat_rule_collection FirewallPolicyRuleCollectionGroup#nat_rule_collection}
	NatRuleCollection interface{} `field:"optional" json:"natRuleCollection" yaml:"natRuleCollection"`
	// network_rule_collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/firewall_policy_rule_collection_group#network_rule_collection FirewallPolicyRuleCollectionGroup#network_rule_collection}
	NetworkRuleCollection interface{} `field:"optional" json:"networkRuleCollection" yaml:"networkRuleCollection"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/firewall_policy_rule_collection_group#timeouts FirewallPolicyRuleCollectionGroup#timeouts}
	Timeouts *FirewallPolicyRuleCollectionGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

