// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsstreaminputblob


type StreamAnalyticsStreamInputBlobSerialization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#type StreamAnalyticsStreamInputBlob#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#encoding StreamAnalyticsStreamInputBlob#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#field_delimiter StreamAnalyticsStreamInputBlob#field_delimiter}.
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
}

