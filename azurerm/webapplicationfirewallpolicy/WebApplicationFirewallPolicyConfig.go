// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebApplicationFirewallPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#location WebApplicationFirewallPolicy#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// managed_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#managed_rules WebApplicationFirewallPolicy#managed_rules}
	ManagedRules *WebApplicationFirewallPolicyManagedRules `field:"required" json:"managedRules" yaml:"managedRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#name WebApplicationFirewallPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#resource_group_name WebApplicationFirewallPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// custom_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#custom_rules WebApplicationFirewallPolicy#custom_rules}
	CustomRules interface{} `field:"optional" json:"customRules" yaml:"customRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#id WebApplicationFirewallPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// policy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#policy_settings WebApplicationFirewallPolicy#policy_settings}
	PolicySettings *WebApplicationFirewallPolicyPolicySettings `field:"optional" json:"policySettings" yaml:"policySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#tags WebApplicationFirewallPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/web_application_firewall_policy#timeouts WebApplicationFirewallPolicy#timeouts}
	Timeouts *WebApplicationFirewallPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

