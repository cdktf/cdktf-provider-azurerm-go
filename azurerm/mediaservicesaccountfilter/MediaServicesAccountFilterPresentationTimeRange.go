// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaservicesaccountfilter


type MediaServicesAccountFilterPresentationTimeRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_services_account_filter#unit_timescale_in_milliseconds MediaServicesAccountFilter#unit_timescale_in_milliseconds}.
	UnitTimescaleInMilliseconds *float64 `field:"required" json:"unitTimescaleInMilliseconds" yaml:"unitTimescaleInMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_services_account_filter#end_in_units MediaServicesAccountFilter#end_in_units}.
	EndInUnits *float64 `field:"optional" json:"endInUnits" yaml:"endInUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_services_account_filter#force_end MediaServicesAccountFilter#force_end}.
	ForceEnd interface{} `field:"optional" json:"forceEnd" yaml:"forceEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_services_account_filter#live_backoff_in_units MediaServicesAccountFilter#live_backoff_in_units}.
	LiveBackoffInUnits *float64 `field:"optional" json:"liveBackoffInUnits" yaml:"liveBackoffInUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_services_account_filter#presentation_window_in_units MediaServicesAccountFilter#presentation_window_in_units}.
	PresentationWindowInUnits *float64 `field:"optional" json:"presentationWindowInUnits" yaml:"presentationWindowInUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_services_account_filter#start_in_units MediaServicesAccountFilter#start_in_units}.
	StartInUnits *float64 `field:"optional" json:"startInUnits" yaml:"startInUnits"`
}

