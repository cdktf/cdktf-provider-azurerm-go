package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCencClearKeyEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#custom_keys_acquisition_url_template MediaStreamingPolicy#custom_keys_acquisition_url_template}.
	CustomKeysAcquisitionUrlTemplate *string `field:"required" json:"customKeysAcquisitionUrlTemplate" yaml:"customKeysAcquisitionUrlTemplate"`
}

