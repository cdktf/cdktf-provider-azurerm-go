// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsreferenceinputblob


type StreamAnalyticsReferenceInputBlobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/stream_analytics_reference_input_blob#create StreamAnalyticsReferenceInputBlob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/stream_analytics_reference_input_blob#delete StreamAnalyticsReferenceInputBlob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/stream_analytics_reference_input_blob#read StreamAnalyticsReferenceInputBlob#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/stream_analytics_reference_input_blob#update StreamAnalyticsReferenceInputBlob#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

