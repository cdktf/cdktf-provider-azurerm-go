// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreaminglocator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaStreamingLocatorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#asset_name MediaStreamingLocator#asset_name}.
	AssetName *string `field:"required" json:"assetName" yaml:"assetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#media_services_account_name MediaStreamingLocator#media_services_account_name}.
	MediaServicesAccountName *string `field:"required" json:"mediaServicesAccountName" yaml:"mediaServicesAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#name MediaStreamingLocator#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#resource_group_name MediaStreamingLocator#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#streaming_policy_name MediaStreamingLocator#streaming_policy_name}.
	StreamingPolicyName *string `field:"required" json:"streamingPolicyName" yaml:"streamingPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#alternative_media_id MediaStreamingLocator#alternative_media_id}.
	AlternativeMediaId *string `field:"optional" json:"alternativeMediaId" yaml:"alternativeMediaId"`
	// content_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#content_key MediaStreamingLocator#content_key}
	ContentKey interface{} `field:"optional" json:"contentKey" yaml:"contentKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#default_content_key_policy_name MediaStreamingLocator#default_content_key_policy_name}.
	DefaultContentKeyPolicyName *string `field:"optional" json:"defaultContentKeyPolicyName" yaml:"defaultContentKeyPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#end_time MediaStreamingLocator#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#filter_names MediaStreamingLocator#filter_names}.
	FilterNames *[]*string `field:"optional" json:"filterNames" yaml:"filterNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#id MediaStreamingLocator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#start_time MediaStreamingLocator#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#streaming_locator_id MediaStreamingLocator#streaming_locator_id}.
	StreamingLocatorId *string `field:"optional" json:"streamingLocatorId" yaml:"streamingLocatorId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_locator#timeouts MediaStreamingLocator#timeouts}
	Timeouts *MediaStreamingLocatorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

