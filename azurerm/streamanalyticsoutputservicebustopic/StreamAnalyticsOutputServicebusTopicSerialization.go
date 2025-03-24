// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputservicebustopic


type StreamAnalyticsOutputServicebusTopicSerialization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_servicebus_topic#type StreamAnalyticsOutputServicebusTopic#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_servicebus_topic#encoding StreamAnalyticsOutputServicebusTopic#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_servicebus_topic#field_delimiter StreamAnalyticsOutputServicebusTopic#field_delimiter}.
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/stream_analytics_output_servicebus_topic#format StreamAnalyticsOutputServicebusTopic#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

