// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendTls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/api_management_backend#validate_certificate_chain ApiManagementBackend#validate_certificate_chain}.
	ValidateCertificateChain interface{} `field:"optional" json:"validateCertificateChain" yaml:"validateCertificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/api_management_backend#validate_certificate_name ApiManagementBackend#validate_certificate_name}.
	ValidateCertificateName interface{} `field:"optional" json:"validateCertificateName" yaml:"validateCertificateName"`
}

