// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputpowerbi

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsOutputPowerbiConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#dataset StreamAnalyticsOutputPowerbi#dataset}.
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#group_id StreamAnalyticsOutputPowerbi#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#group_name StreamAnalyticsOutputPowerbi#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#name StreamAnalyticsOutputPowerbi#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#stream_analytics_job_id StreamAnalyticsOutputPowerbi#stream_analytics_job_id}.
	StreamAnalyticsJobId *string `field:"required" json:"streamAnalyticsJobId" yaml:"streamAnalyticsJobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#table StreamAnalyticsOutputPowerbi#table}.
	Table *string `field:"required" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#id StreamAnalyticsOutputPowerbi#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#timeouts StreamAnalyticsOutputPowerbi#timeouts}
	Timeouts *StreamAnalyticsOutputPowerbiTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#token_user_display_name StreamAnalyticsOutputPowerbi#token_user_display_name}.
	TokenUserDisplayName *string `field:"optional" json:"tokenUserDisplayName" yaml:"tokenUserDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_powerbi#token_user_principal_name StreamAnalyticsOutputPowerbi#token_user_principal_name}.
	TokenUserPrincipalName *string `field:"optional" json:"tokenUserPrincipalName" yaml:"tokenUserPrincipalName"`
}

