// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorcustomhttpsconfiguration


type FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/frontdoor_custom_https_configuration#azure_key_vault_certificate_secret_name FrontdoorCustomHttpsConfiguration#azure_key_vault_certificate_secret_name}.
	AzureKeyVaultCertificateSecretName *string `field:"optional" json:"azureKeyVaultCertificateSecretName" yaml:"azureKeyVaultCertificateSecretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/frontdoor_custom_https_configuration#azure_key_vault_certificate_secret_version FrontdoorCustomHttpsConfiguration#azure_key_vault_certificate_secret_version}.
	AzureKeyVaultCertificateSecretVersion *string `field:"optional" json:"azureKeyVaultCertificateSecretVersion" yaml:"azureKeyVaultCertificateSecretVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/frontdoor_custom_https_configuration#azure_key_vault_certificate_vault_id FrontdoorCustomHttpsConfiguration#azure_key_vault_certificate_vault_id}.
	AzureKeyVaultCertificateVaultId *string `field:"optional" json:"azureKeyVaultCertificateVaultId" yaml:"azureKeyVaultCertificateVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/frontdoor_custom_https_configuration#certificate_source FrontdoorCustomHttpsConfiguration#certificate_source}.
	CertificateSource *string `field:"optional" json:"certificateSource" yaml:"certificateSource"`
}

