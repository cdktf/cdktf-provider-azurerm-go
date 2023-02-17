package cognitiveaccount


type CognitiveAccountCustomerManagedKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_account#key_vault_key_id CognitiveAccount#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_account#identity_client_id CognitiveAccount#identity_client_id}.
	IdentityClientId *string `field:"optional" json:"identityClientId" yaml:"identityClientId"`
}

