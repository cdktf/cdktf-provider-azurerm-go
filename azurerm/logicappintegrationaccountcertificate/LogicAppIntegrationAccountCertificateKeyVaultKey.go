package logicappintegrationaccountcertificate


type LogicAppIntegrationAccountCertificateKeyVaultKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_certificate#key_name LogicAppIntegrationAccountCertificate#key_name}.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_certificate#key_vault_id LogicAppIntegrationAccountCertificate#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_certificate#key_version LogicAppIntegrationAccountCertificate#key_version}.
	KeyVersion *string `field:"optional" json:"keyVersion" yaml:"keyVersion"`
}

