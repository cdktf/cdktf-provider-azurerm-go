// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCencDrmPlayready struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_streaming_policy#custom_attributes MediaStreamingPolicy#custom_attributes}.
	CustomAttributes *string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/media_streaming_policy#custom_license_acquisition_url_template MediaStreamingPolicy#custom_license_acquisition_url_template}.
	CustomLicenseAcquisitionUrlTemplate *string `field:"optional" json:"customLicenseAcquisitionUrlTemplate" yaml:"customLicenseAcquisitionUrlTemplate"`
}

