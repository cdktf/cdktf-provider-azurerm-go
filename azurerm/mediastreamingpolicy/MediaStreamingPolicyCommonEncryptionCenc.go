// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCenc struct {
	// clear_key_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/media_streaming_policy#clear_key_encryption MediaStreamingPolicy#clear_key_encryption}
	ClearKeyEncryption *MediaStreamingPolicyCommonEncryptionCencClearKeyEncryption `field:"optional" json:"clearKeyEncryption" yaml:"clearKeyEncryption"`
	// clear_track block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/media_streaming_policy#clear_track MediaStreamingPolicy#clear_track}
	ClearTrack interface{} `field:"optional" json:"clearTrack" yaml:"clearTrack"`
	// content_key_to_track_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/media_streaming_policy#content_key_to_track_mapping MediaStreamingPolicy#content_key_to_track_mapping}
	ContentKeyToTrackMapping interface{} `field:"optional" json:"contentKeyToTrackMapping" yaml:"contentKeyToTrackMapping"`
	// default_content_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/media_streaming_policy#default_content_key MediaStreamingPolicy#default_content_key}
	DefaultContentKey *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey `field:"optional" json:"defaultContentKey" yaml:"defaultContentKey"`
	// drm_playready block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/media_streaming_policy#drm_playready MediaStreamingPolicy#drm_playready}
	DrmPlayready *MediaStreamingPolicyCommonEncryptionCencDrmPlayready `field:"optional" json:"drmPlayready" yaml:"drmPlayready"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/media_streaming_policy#drm_widevine_custom_license_acquisition_url_template MediaStreamingPolicy#drm_widevine_custom_license_acquisition_url_template}.
	DrmWidevineCustomLicenseAcquisitionUrlTemplate *string `field:"optional" json:"drmWidevineCustomLicenseAcquisitionUrlTemplate" yaml:"drmWidevineCustomLicenseAcquisitionUrlTemplate"`
	// enabled_protocols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/media_streaming_policy#enabled_protocols MediaStreamingPolicy#enabled_protocols}
	EnabledProtocols *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols `field:"optional" json:"enabledProtocols" yaml:"enabledProtocols"`
}

