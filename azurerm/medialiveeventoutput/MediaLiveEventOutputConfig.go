// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveeventoutput

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaLiveEventOutputConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#archive_window_duration MediaLiveEventOutput#archive_window_duration}.
	ArchiveWindowDuration *string `field:"required" json:"archiveWindowDuration" yaml:"archiveWindowDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#asset_name MediaLiveEventOutput#asset_name}.
	AssetName *string `field:"required" json:"assetName" yaml:"assetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#live_event_id MediaLiveEventOutput#live_event_id}.
	LiveEventId *string `field:"required" json:"liveEventId" yaml:"liveEventId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#name MediaLiveEventOutput#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#description MediaLiveEventOutput#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#hls_fragments_per_ts_segment MediaLiveEventOutput#hls_fragments_per_ts_segment}.
	HlsFragmentsPerTsSegment *float64 `field:"optional" json:"hlsFragmentsPerTsSegment" yaml:"hlsFragmentsPerTsSegment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#id MediaLiveEventOutput#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#manifest_name MediaLiveEventOutput#manifest_name}.
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#output_snap_time_in_seconds MediaLiveEventOutput#output_snap_time_in_seconds}.
	OutputSnapTimeInSeconds *float64 `field:"optional" json:"outputSnapTimeInSeconds" yaml:"outputSnapTimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#rewind_window_duration MediaLiveEventOutput#rewind_window_duration}.
	RewindWindowDuration *string `field:"optional" json:"rewindWindowDuration" yaml:"rewindWindowDuration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event_output#timeouts MediaLiveEventOutput#timeouts}
	Timeouts *MediaLiveEventOutputTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

