// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputservicebusqueue


type StreamAnalyticsOutputServicebusQueueSerialization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/stream_analytics_output_servicebus_queue#type StreamAnalyticsOutputServicebusQueue#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/stream_analytics_output_servicebus_queue#encoding StreamAnalyticsOutputServicebusQueue#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/stream_analytics_output_servicebus_queue#field_delimiter StreamAnalyticsOutputServicebusQueue#field_delimiter}.
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/stream_analytics_output_servicebus_queue#format StreamAnalyticsOutputServicebusQueue#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

