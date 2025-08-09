// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputservicebusqueue


type StreamAnalyticsOutputServicebusQueueTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/stream_analytics_output_servicebus_queue#create StreamAnalyticsOutputServicebusQueue#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/stream_analytics_output_servicebus_queue#delete StreamAnalyticsOutputServicebusQueue#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/stream_analytics_output_servicebus_queue#read StreamAnalyticsOutputServicebusQueue#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/stream_analytics_output_servicebus_queue#update StreamAnalyticsOutputServicebusQueue#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

