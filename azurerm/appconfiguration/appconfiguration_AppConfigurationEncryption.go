package appconfiguration


type AppConfigurationEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration#identity_client_id AppConfiguration#identity_client_id}.
	IdentityClientId *string `field:"optional" json:"identityClientId" yaml:"identityClientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_configuration#key_vault_key_identifier AppConfiguration#key_vault_key_identifier}.
	KeyVaultKeyIdentifier *string `field:"optional" json:"keyVaultKeyIdentifier" yaml:"keyVaultKeyIdentifier"`
}
