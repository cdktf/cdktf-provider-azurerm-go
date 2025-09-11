// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#name CdnFrontdoorProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#resource_group_name CdnFrontdoorProfile#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#sku_name CdnFrontdoorProfile#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#id CdnFrontdoorProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#identity CdnFrontdoorProfile#identity}
	Identity *CdnFrontdoorProfileIdentity `field:"optional" json:"identity" yaml:"identity"`
	// log_scrubbing_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#log_scrubbing_rule CdnFrontdoorProfile#log_scrubbing_rule}
	LogScrubbingRule interface{} `field:"optional" json:"logScrubbingRule" yaml:"logScrubbingRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#response_timeout_seconds CdnFrontdoorProfile#response_timeout_seconds}.
	ResponseTimeoutSeconds *float64 `field:"optional" json:"responseTimeoutSeconds" yaml:"responseTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#tags CdnFrontdoorProfile#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_profile#timeouts CdnFrontdoorProfile#timeouts}
	Timeouts *CdnFrontdoorProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

