// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputsynapse


type StreamAnalyticsOutputSynapseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/stream_analytics_output_synapse#create StreamAnalyticsOutputSynapse#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/stream_analytics_output_synapse#delete StreamAnalyticsOutputSynapse#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/stream_analytics_output_synapse#read StreamAnalyticsOutputSynapse#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/stream_analytics_output_synapse#update StreamAnalyticsOutputSynapse#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

