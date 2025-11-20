// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorfirewallpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorFirewallPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#mode CdnFrontdoorFirewallPolicy#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#name CdnFrontdoorFirewallPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#resource_group_name CdnFrontdoorFirewallPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#sku_name CdnFrontdoorFirewallPolicy#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#captcha_cookie_expiration_in_minutes CdnFrontdoorFirewallPolicy#captcha_cookie_expiration_in_minutes}.
	CaptchaCookieExpirationInMinutes *float64 `field:"optional" json:"captchaCookieExpirationInMinutes" yaml:"captchaCookieExpirationInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#custom_block_response_body CdnFrontdoorFirewallPolicy#custom_block_response_body}.
	CustomBlockResponseBody *string `field:"optional" json:"customBlockResponseBody" yaml:"customBlockResponseBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#custom_block_response_status_code CdnFrontdoorFirewallPolicy#custom_block_response_status_code}.
	CustomBlockResponseStatusCode *float64 `field:"optional" json:"customBlockResponseStatusCode" yaml:"customBlockResponseStatusCode"`
	// custom_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#custom_rule CdnFrontdoorFirewallPolicy#custom_rule}
	CustomRule interface{} `field:"optional" json:"customRule" yaml:"customRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#enabled CdnFrontdoorFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#id CdnFrontdoorFirewallPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#js_challenge_cookie_expiration_in_minutes CdnFrontdoorFirewallPolicy#js_challenge_cookie_expiration_in_minutes}.
	JsChallengeCookieExpirationInMinutes *float64 `field:"optional" json:"jsChallengeCookieExpirationInMinutes" yaml:"jsChallengeCookieExpirationInMinutes"`
	// log_scrubbing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#log_scrubbing CdnFrontdoorFirewallPolicy#log_scrubbing}
	LogScrubbing *CdnFrontdoorFirewallPolicyLogScrubbing `field:"optional" json:"logScrubbing" yaml:"logScrubbing"`
	// managed_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#managed_rule CdnFrontdoorFirewallPolicy#managed_rule}
	ManagedRule interface{} `field:"optional" json:"managedRule" yaml:"managedRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#redirect_url CdnFrontdoorFirewallPolicy#redirect_url}.
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#request_body_check_enabled CdnFrontdoorFirewallPolicy#request_body_check_enabled}.
	RequestBodyCheckEnabled interface{} `field:"optional" json:"requestBodyCheckEnabled" yaml:"requestBodyCheckEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#tags CdnFrontdoorFirewallPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/cdn_frontdoor_firewall_policy#timeouts CdnFrontdoorFirewallPolicy#timeouts}
	Timeouts *CdnFrontdoorFirewallPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

