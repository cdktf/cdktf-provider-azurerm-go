package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCenc struct {
	// default_content_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#default_content_key MediaStreamingPolicy#default_content_key}
	DefaultContentKey *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey `field:"optional" json:"defaultContentKey" yaml:"defaultContentKey"`
	// drm_playready block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#drm_playready MediaStreamingPolicy#drm_playready}
	DrmPlayready *MediaStreamingPolicyCommonEncryptionCencDrmPlayready `field:"optional" json:"drmPlayready" yaml:"drmPlayready"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#drm_widevine_custom_license_acquisition_url_template MediaStreamingPolicy#drm_widevine_custom_license_acquisition_url_template}.
	DrmWidevineCustomLicenseAcquisitionUrlTemplate *string `field:"optional" json:"drmWidevineCustomLicenseAcquisitionUrlTemplate" yaml:"drmWidevineCustomLicenseAcquisitionUrlTemplate"`
	// enabled_protocols block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#enabled_protocols MediaStreamingPolicy#enabled_protocols}
	EnabledProtocols *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols `field:"optional" json:"enabledProtocols" yaml:"enabledProtocols"`
}

