// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementHostnameConfigurationProxy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/api_management#host_name ApiManagement#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/api_management#certificate ApiManagement#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/api_management#certificate_password ApiManagement#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/api_management#default_ssl_binding ApiManagement#default_ssl_binding}.
	DefaultSslBinding interface{} `field:"optional" json:"defaultSslBinding" yaml:"defaultSslBinding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/api_management#key_vault_id ApiManagement#key_vault_id}.
	KeyVaultId *string `field:"optional" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/api_management#negotiate_client_certificate ApiManagement#negotiate_client_certificate}.
	NegotiateClientCertificate interface{} `field:"optional" json:"negotiateClientCertificate" yaml:"negotiateClientCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/api_management#ssl_keyvault_identity_client_id ApiManagement#ssl_keyvault_identity_client_id}.
	SslKeyvaultIdentityClientId *string `field:"optional" json:"sslKeyvaultIdentityClientId" yaml:"sslKeyvaultIdentityClientId"`
}

