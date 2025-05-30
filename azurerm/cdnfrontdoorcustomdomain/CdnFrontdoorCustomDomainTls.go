// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorcustomdomain


type CdnFrontdoorCustomDomainTls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/cdn_frontdoor_custom_domain#cdn_frontdoor_secret_id CdnFrontdoorCustomDomain#cdn_frontdoor_secret_id}.
	CdnFrontdoorSecretId *string `field:"optional" json:"cdnFrontdoorSecretId" yaml:"cdnFrontdoorSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/cdn_frontdoor_custom_domain#certificate_type CdnFrontdoorCustomDomain#certificate_type}.
	CertificateType *string `field:"optional" json:"certificateType" yaml:"certificateType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/cdn_frontdoor_custom_domain#minimum_tls_version CdnFrontdoorCustomDomain#minimum_tls_version}.
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
}

