// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsfunctionjavascriptudf


type StreamAnalyticsFunctionJavascriptUdfInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/stream_analytics_function_javascript_udf#type StreamAnalyticsFunctionJavascriptUdf#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/stream_analytics_function_javascript_udf#configuration_parameter StreamAnalyticsFunctionJavascriptUdf#configuration_parameter}.
	ConfigurationParameter interface{} `field:"optional" json:"configurationParameter" yaml:"configurationParameter"`
}

