// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorsecret


type CdnFrontdoorSecretSecretCustomerCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cdn_frontdoor_secret#key_vault_certificate_id CdnFrontdoorSecret#key_vault_certificate_id}.
	KeyVaultCertificateId *string `field:"required" json:"keyVaultCertificateId" yaml:"keyVaultCertificateId"`
}

