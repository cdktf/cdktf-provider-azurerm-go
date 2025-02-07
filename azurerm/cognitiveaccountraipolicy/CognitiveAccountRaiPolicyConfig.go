// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitiveaccountraipolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitiveAccountRaiPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cognitive_account_rai_policy#base_policy_name CognitiveAccountRaiPolicy#base_policy_name}.
	BasePolicyName *string `field:"required" json:"basePolicyName" yaml:"basePolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cognitive_account_rai_policy#cognitive_account_id CognitiveAccountRaiPolicy#cognitive_account_id}.
	CognitiveAccountId *string `field:"required" json:"cognitiveAccountId" yaml:"cognitiveAccountId"`
	// content_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cognitive_account_rai_policy#content_filter CognitiveAccountRaiPolicy#content_filter}
	ContentFilter interface{} `field:"required" json:"contentFilter" yaml:"contentFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cognitive_account_rai_policy#name CognitiveAccountRaiPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cognitive_account_rai_policy#id CognitiveAccountRaiPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cognitive_account_rai_policy#mode CognitiveAccountRaiPolicy#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cognitive_account_rai_policy#tags CognitiveAccountRaiPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cognitive_account_rai_policy#timeouts CognitiveAccountRaiPolicy#timeouts}
	Timeouts *CognitiveAccountRaiPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

