// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputpowerbi


type StreamAnalyticsOutputPowerbiTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/stream_analytics_output_powerbi#create StreamAnalyticsOutputPowerbi#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/stream_analytics_output_powerbi#delete StreamAnalyticsOutputPowerbi#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/stream_analytics_output_powerbi#read StreamAnalyticsOutputPowerbi#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/stream_analytics_output_powerbi#update StreamAnalyticsOutputPowerbi#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

