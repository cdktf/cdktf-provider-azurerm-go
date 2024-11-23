// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/api_management#encoded_certificate ApiManagement#encoded_certificate}.
	EncodedCertificate *string `field:"required" json:"encodedCertificate" yaml:"encodedCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/api_management#store_name ApiManagement#store_name}.
	StoreName *string `field:"required" json:"storeName" yaml:"storeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/api_management#certificate_password ApiManagement#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
}

