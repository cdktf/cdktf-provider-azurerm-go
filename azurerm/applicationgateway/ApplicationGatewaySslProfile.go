// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewaySslProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// ssl_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/application_gateway#ssl_policy ApplicationGateway#ssl_policy}
	SslPolicy *ApplicationGatewaySslProfileSslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/application_gateway#trusted_client_certificate_names ApplicationGateway#trusted_client_certificate_names}.
	TrustedClientCertificateNames *[]*string `field:"optional" json:"trustedClientCertificateNames" yaml:"trustedClientCertificateNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/application_gateway#verify_client_certificate_revocation ApplicationGateway#verify_client_certificate_revocation}.
	VerifyClientCertificateRevocation *string `field:"optional" json:"verifyClientCertificateRevocation" yaml:"verifyClientCertificateRevocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/application_gateway#verify_client_cert_issuer_dn ApplicationGateway#verify_client_cert_issuer_dn}.
	VerifyClientCertIssuerDn interface{} `field:"optional" json:"verifyClientCertIssuerDn" yaml:"verifyClientCertIssuerDn"`
}

