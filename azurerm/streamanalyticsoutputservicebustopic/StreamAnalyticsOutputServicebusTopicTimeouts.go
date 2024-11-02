// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputservicebustopic


type StreamAnalyticsOutputServicebusTopicTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/stream_analytics_output_servicebus_topic#create StreamAnalyticsOutputServicebusTopic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/stream_analytics_output_servicebus_topic#delete StreamAnalyticsOutputServicebusTopic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/stream_analytics_output_servicebus_topic#read StreamAnalyticsOutputServicebusTopic#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/stream_analytics_output_servicebus_topic#update StreamAnalyticsOutputServicebusTopic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

