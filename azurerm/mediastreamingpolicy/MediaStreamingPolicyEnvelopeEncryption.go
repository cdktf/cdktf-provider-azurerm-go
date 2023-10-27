// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingpolicy


type MediaStreamingPolicyEnvelopeEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/media_streaming_policy#custom_keys_acquisition_url_template MediaStreamingPolicy#custom_keys_acquisition_url_template}.
	CustomKeysAcquisitionUrlTemplate *string `field:"optional" json:"customKeysAcquisitionUrlTemplate" yaml:"customKeysAcquisitionUrlTemplate"`
	// default_content_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/media_streaming_policy#default_content_key MediaStreamingPolicy#default_content_key}
	DefaultContentKey *MediaStreamingPolicyEnvelopeEncryptionDefaultContentKey `field:"optional" json:"defaultContentKey" yaml:"defaultContentKey"`
	// enabled_protocols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/media_streaming_policy#enabled_protocols MediaStreamingPolicy#enabled_protocols}
	EnabledProtocols *MediaStreamingPolicyEnvelopeEncryptionEnabledProtocols `field:"optional" json:"enabledProtocols" yaml:"enabledProtocols"`
}

