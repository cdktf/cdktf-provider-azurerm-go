// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveevent


type MediaLiveEventInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/media_live_event#access_token MediaLiveEvent#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// ip_access_control_allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/media_live_event#ip_access_control_allow MediaLiveEvent#ip_access_control_allow}
	IpAccessControlAllow interface{} `field:"optional" json:"ipAccessControlAllow" yaml:"ipAccessControlAllow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/media_live_event#key_frame_interval_duration MediaLiveEvent#key_frame_interval_duration}.
	KeyFrameIntervalDuration *string `field:"optional" json:"keyFrameIntervalDuration" yaml:"keyFrameIntervalDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/media_live_event#streaming_protocol MediaLiveEvent#streaming_protocol}.
	StreamingProtocol *string `field:"optional" json:"streamingProtocol" yaml:"streamingProtocol"`
}

