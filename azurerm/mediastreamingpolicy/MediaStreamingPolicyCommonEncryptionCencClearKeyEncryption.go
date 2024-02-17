// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCencClearKeyEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/media_streaming_policy#custom_keys_acquisition_url_template MediaStreamingPolicy#custom_keys_acquisition_url_template}.
	CustomKeysAcquisitionUrlTemplate *string `field:"required" json:"customKeysAcquisitionUrlTemplate" yaml:"customKeysAcquisitionUrlTemplate"`
}

