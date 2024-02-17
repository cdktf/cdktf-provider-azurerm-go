// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputblob


type StreamAnalyticsOutputBlobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/stream_analytics_output_blob#create StreamAnalyticsOutputBlob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/stream_analytics_output_blob#delete StreamAnalyticsOutputBlob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/stream_analytics_output_blob#read StreamAnalyticsOutputBlob#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/stream_analytics_output_blob#update StreamAnalyticsOutputBlob#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

