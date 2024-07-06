// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottimeseriesinsightseventsourceiothub


type IotTimeSeriesInsightsEventSourceIothubTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/iot_time_series_insights_event_source_iothub#create IotTimeSeriesInsightsEventSourceIothub#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/iot_time_series_insights_event_source_iothub#delete IotTimeSeriesInsightsEventSourceIothub#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/iot_time_series_insights_event_source_iothub#read IotTimeSeriesInsightsEventSourceIothub#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/iot_time_series_insights_event_source_iothub#update IotTimeSeriesInsightsEventSourceIothub#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

