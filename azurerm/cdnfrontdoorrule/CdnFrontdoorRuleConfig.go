// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorRuleConfig struct {
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
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_rule#actions CdnFrontdoorRule#actions}
	Actions *CdnFrontdoorRuleActions `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_rule#cdn_frontdoor_rule_set_id CdnFrontdoorRule#cdn_frontdoor_rule_set_id}.
	CdnFrontdoorRuleSetId *string `field:"required" json:"cdnFrontdoorRuleSetId" yaml:"cdnFrontdoorRuleSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_rule#name CdnFrontdoorRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_rule#order CdnFrontdoorRule#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_rule#behavior_on_match CdnFrontdoorRule#behavior_on_match}.
	BehaviorOnMatch *string `field:"optional" json:"behaviorOnMatch" yaml:"behaviorOnMatch"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_rule#conditions CdnFrontdoorRule#conditions}
	Conditions *CdnFrontdoorRuleConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_rule#id CdnFrontdoorRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/cdn_frontdoor_rule#timeouts CdnFrontdoorRule#timeouts}
	Timeouts *CdnFrontdoorRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

