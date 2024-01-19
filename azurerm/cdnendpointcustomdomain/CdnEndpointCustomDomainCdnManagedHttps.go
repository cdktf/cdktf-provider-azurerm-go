// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpointcustomdomain


type CdnEndpointCustomDomainCdnManagedHttps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/cdn_endpoint_custom_domain#certificate_type CdnEndpointCustomDomain#certificate_type}.
	CertificateType *string `field:"required" json:"certificateType" yaml:"certificateType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/cdn_endpoint_custom_domain#protocol_type CdnEndpointCustomDomain#protocol_type}.
	ProtocolType *string `field:"required" json:"protocolType" yaml:"protocolType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/cdn_endpoint_custom_domain#tls_version CdnEndpointCustomDomain#tls_version}.
	TlsVersion *string `field:"optional" json:"tlsVersion" yaml:"tlsVersion"`
}

