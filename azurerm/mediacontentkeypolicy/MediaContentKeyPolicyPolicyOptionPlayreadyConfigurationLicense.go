// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediacontentkeypolicy


type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicense struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#allow_test_devices MediaContentKeyPolicy#allow_test_devices}.
	AllowTestDevices interface{} `field:"optional" json:"allowTestDevices" yaml:"allowTestDevices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#begin_date MediaContentKeyPolicy#begin_date}.
	BeginDate *string `field:"optional" json:"beginDate" yaml:"beginDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#content_key_location_from_header_enabled MediaContentKeyPolicy#content_key_location_from_header_enabled}.
	ContentKeyLocationFromHeaderEnabled interface{} `field:"optional" json:"contentKeyLocationFromHeaderEnabled" yaml:"contentKeyLocationFromHeaderEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#content_key_location_from_key_id MediaContentKeyPolicy#content_key_location_from_key_id}.
	ContentKeyLocationFromKeyId *string `field:"optional" json:"contentKeyLocationFromKeyId" yaml:"contentKeyLocationFromKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#content_type MediaContentKeyPolicy#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#expiration_date MediaContentKeyPolicy#expiration_date}.
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#grace_period MediaContentKeyPolicy#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#license_type MediaContentKeyPolicy#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// play_right block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#play_right MediaContentKeyPolicy#play_right}
	PlayRight *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight `field:"optional" json:"playRight" yaml:"playRight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#relative_begin_date MediaContentKeyPolicy#relative_begin_date}.
	RelativeBeginDate *string `field:"optional" json:"relativeBeginDate" yaml:"relativeBeginDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#relative_expiration_date MediaContentKeyPolicy#relative_expiration_date}.
	RelativeExpirationDate *string `field:"optional" json:"relativeExpirationDate" yaml:"relativeExpirationDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/media_content_key_policy#security_level MediaContentKeyPolicy#security_level}.
	SecurityLevel *string `field:"optional" json:"securityLevel" yaml:"securityLevel"`
}

