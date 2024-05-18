// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveevent


type MediaLiveEventPreview struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_live_event#alternative_media_id MediaLiveEvent#alternative_media_id}.
	AlternativeMediaId *string `field:"optional" json:"alternativeMediaId" yaml:"alternativeMediaId"`
	// ip_access_control_allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_live_event#ip_access_control_allow MediaLiveEvent#ip_access_control_allow}
	IpAccessControlAllow interface{} `field:"optional" json:"ipAccessControlAllow" yaml:"ipAccessControlAllow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_live_event#preview_locator MediaLiveEvent#preview_locator}.
	PreviewLocator *string `field:"optional" json:"previewLocator" yaml:"previewLocator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_live_event#streaming_policy_name MediaLiveEvent#streaming_policy_name}.
	StreamingPolicyName *string `field:"optional" json:"streamingPolicyName" yaml:"streamingPolicyName"`
}

