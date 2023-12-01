// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/media_streaming_policy#allow_persistent_license MediaStreamingPolicy#allow_persistent_license}.
	AllowPersistentLicense interface{} `field:"optional" json:"allowPersistentLicense" yaml:"allowPersistentLicense"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/media_streaming_policy#custom_license_acquisition_url_template MediaStreamingPolicy#custom_license_acquisition_url_template}.
	CustomLicenseAcquisitionUrlTemplate *string `field:"optional" json:"customLicenseAcquisitionUrlTemplate" yaml:"customLicenseAcquisitionUrlTemplate"`
}

