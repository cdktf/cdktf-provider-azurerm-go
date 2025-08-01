// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputeventhub


type StreamAnalyticsOutputEventhubTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/stream_analytics_output_eventhub#create StreamAnalyticsOutputEventhub#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/stream_analytics_output_eventhub#delete StreamAnalyticsOutputEventhub#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/stream_analytics_output_eventhub#read StreamAnalyticsOutputEventhub#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/stream_analytics_output_eventhub#update StreamAnalyticsOutputEventhub#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

