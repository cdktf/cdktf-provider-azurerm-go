// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsfunctionjavascriptuda

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsFunctionJavascriptUdaConfig struct {
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
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/stream_analytics_function_javascript_uda#input StreamAnalyticsFunctionJavascriptUda#input}
	Input interface{} `field:"required" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/stream_analytics_function_javascript_uda#name StreamAnalyticsFunctionJavascriptUda#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/stream_analytics_function_javascript_uda#output StreamAnalyticsFunctionJavascriptUda#output}
	Output *StreamAnalyticsFunctionJavascriptUdaOutput `field:"required" json:"output" yaml:"output"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/stream_analytics_function_javascript_uda#script StreamAnalyticsFunctionJavascriptUda#script}.
	Script *string `field:"required" json:"script" yaml:"script"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/stream_analytics_function_javascript_uda#stream_analytics_job_id StreamAnalyticsFunctionJavascriptUda#stream_analytics_job_id}.
	StreamAnalyticsJobId *string `field:"required" json:"streamAnalyticsJobId" yaml:"streamAnalyticsJobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/stream_analytics_function_javascript_uda#id StreamAnalyticsFunctionJavascriptUda#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/stream_analytics_function_javascript_uda#timeouts StreamAnalyticsFunctionJavascriptUda#timeouts}
	Timeouts *StreamAnalyticsFunctionJavascriptUdaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

