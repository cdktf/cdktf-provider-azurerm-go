package keyvaultkey


type KeyVaultKeyRotationPolicyAutomatic struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_key#time_after_creation KeyVaultKey#time_after_creation}.
	TimeAfterCreation *string `field:"optional" json:"timeAfterCreation" yaml:"timeAfterCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_key#time_before_expiry KeyVaultKey#time_before_expiry}.
	TimeBeforeExpiry *string `field:"optional" json:"timeBeforeExpiry" yaml:"timeBeforeExpiry"`
}

