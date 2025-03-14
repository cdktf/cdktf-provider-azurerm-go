// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorfirewallpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrontdoorFirewallPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#name FrontdoorFirewallPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#resource_group_name FrontdoorFirewallPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#custom_block_response_body FrontdoorFirewallPolicy#custom_block_response_body}.
	CustomBlockResponseBody *string `field:"optional" json:"customBlockResponseBody" yaml:"customBlockResponseBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#custom_block_response_status_code FrontdoorFirewallPolicy#custom_block_response_status_code}.
	CustomBlockResponseStatusCode *float64 `field:"optional" json:"customBlockResponseStatusCode" yaml:"customBlockResponseStatusCode"`
	// custom_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#custom_rule FrontdoorFirewallPolicy#custom_rule}
	CustomRule interface{} `field:"optional" json:"customRule" yaml:"customRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#enabled FrontdoorFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#id FrontdoorFirewallPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// managed_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#managed_rule FrontdoorFirewallPolicy#managed_rule}
	ManagedRule interface{} `field:"optional" json:"managedRule" yaml:"managedRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#mode FrontdoorFirewallPolicy#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#redirect_url FrontdoorFirewallPolicy#redirect_url}.
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#tags FrontdoorFirewallPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/frontdoor_firewall_policy#timeouts FrontdoorFirewallPolicy#timeouts}
	Timeouts *FrontdoorFirewallPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

