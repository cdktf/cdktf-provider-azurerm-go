// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementcustomdomain


type ApiManagementCustomDomainGateway struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/api_management_custom_domain#host_name ApiManagementCustomDomain#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/api_management_custom_domain#certificate ApiManagementCustomDomain#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/api_management_custom_domain#certificate_password ApiManagementCustomDomain#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/api_management_custom_domain#default_ssl_binding ApiManagementCustomDomain#default_ssl_binding}.
	DefaultSslBinding interface{} `field:"optional" json:"defaultSslBinding" yaml:"defaultSslBinding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/api_management_custom_domain#key_vault_certificate_id ApiManagementCustomDomain#key_vault_certificate_id}.
	KeyVaultCertificateId *string `field:"optional" json:"keyVaultCertificateId" yaml:"keyVaultCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/api_management_custom_domain#key_vault_id ApiManagementCustomDomain#key_vault_id}.
	KeyVaultId *string `field:"optional" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/api_management_custom_domain#negotiate_client_certificate ApiManagementCustomDomain#negotiate_client_certificate}.
	NegotiateClientCertificate interface{} `field:"optional" json:"negotiateClientCertificate" yaml:"negotiateClientCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/api_management_custom_domain#ssl_keyvault_identity_client_id ApiManagementCustomDomain#ssl_keyvault_identity_client_id}.
	SslKeyvaultIdentityClientId *string `field:"optional" json:"sslKeyvaultIdentityClientId" yaml:"sslKeyvaultIdentityClientId"`
}

