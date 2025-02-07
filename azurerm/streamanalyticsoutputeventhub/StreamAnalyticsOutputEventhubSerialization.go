// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputeventhub


type StreamAnalyticsOutputEventhubSerialization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/stream_analytics_output_eventhub#type StreamAnalyticsOutputEventhub#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/stream_analytics_output_eventhub#encoding StreamAnalyticsOutputEventhub#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/stream_analytics_output_eventhub#field_delimiter StreamAnalyticsOutputEventhub#field_delimiter}.
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/stream_analytics_output_eventhub#format StreamAnalyticsOutputEventhub#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

