package containerregistry


type ContainerRegistryEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#enabled ContainerRegistry#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#identity_client_id ContainerRegistry#identity_client_id}.
	IdentityClientId *string `field:"optional" json:"identityClientId" yaml:"identityClientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#key_vault_key_id ContainerRegistry#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
}
