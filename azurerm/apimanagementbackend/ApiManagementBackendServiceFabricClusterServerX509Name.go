// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendServiceFabricClusterServerX509Name struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_backend#issuer_certificate_thumbprint ApiManagementBackend#issuer_certificate_thumbprint}.
	IssuerCertificateThumbprint *string `field:"required" json:"issuerCertificateThumbprint" yaml:"issuerCertificateThumbprint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_backend#name ApiManagementBackend#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

