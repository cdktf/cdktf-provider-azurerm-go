// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaassetfilter


type MediaAssetFilterPresentationTimeRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_asset_filter#end_in_units MediaAssetFilter#end_in_units}.
	EndInUnits *float64 `field:"optional" json:"endInUnits" yaml:"endInUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_asset_filter#force_end MediaAssetFilter#force_end}.
	ForceEnd interface{} `field:"optional" json:"forceEnd" yaml:"forceEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_asset_filter#live_backoff_in_units MediaAssetFilter#live_backoff_in_units}.
	LiveBackoffInUnits *float64 `field:"optional" json:"liveBackoffInUnits" yaml:"liveBackoffInUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_asset_filter#presentation_window_in_units MediaAssetFilter#presentation_window_in_units}.
	PresentationWindowInUnits *float64 `field:"optional" json:"presentationWindowInUnits" yaml:"presentationWindowInUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_asset_filter#start_in_units MediaAssetFilter#start_in_units}.
	StartInUnits *float64 `field:"optional" json:"startInUnits" yaml:"startInUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_asset_filter#unit_timescale_in_miliseconds MediaAssetFilter#unit_timescale_in_miliseconds}.
	UnitTimescaleInMiliseconds *float64 `field:"optional" json:"unitTimescaleInMiliseconds" yaml:"unitTimescaleInMiliseconds"`
}

