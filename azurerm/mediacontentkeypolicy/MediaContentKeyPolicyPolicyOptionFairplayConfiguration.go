// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediacontentkeypolicy


type MediaContentKeyPolicyPolicyOptionFairplayConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_content_key_policy#ask MediaContentKeyPolicy#ask}.
	Ask *string `field:"optional" json:"ask" yaml:"ask"`
	// offline_rental_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_content_key_policy#offline_rental_configuration MediaContentKeyPolicy#offline_rental_configuration}
	OfflineRentalConfiguration *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration `field:"optional" json:"offlineRentalConfiguration" yaml:"offlineRentalConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_content_key_policy#pfx MediaContentKeyPolicy#pfx}.
	Pfx *string `field:"optional" json:"pfx" yaml:"pfx"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_content_key_policy#pfx_password MediaContentKeyPolicy#pfx_password}.
	PfxPassword *string `field:"optional" json:"pfxPassword" yaml:"pfxPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_content_key_policy#rental_and_lease_key_type MediaContentKeyPolicy#rental_and_lease_key_type}.
	RentalAndLeaseKeyType *string `field:"optional" json:"rentalAndLeaseKeyType" yaml:"rentalAndLeaseKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_content_key_policy#rental_duration_seconds MediaContentKeyPolicy#rental_duration_seconds}.
	RentalDurationSeconds *float64 `field:"optional" json:"rentalDurationSeconds" yaml:"rentalDurationSeconds"`
}

