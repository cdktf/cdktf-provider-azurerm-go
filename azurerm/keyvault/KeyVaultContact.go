package keyvault


type KeyVaultContact struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault#email KeyVault#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault#name KeyVault#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault#phone KeyVault#phone}.
	Phone *string `field:"optional" json:"phone" yaml:"phone"`
}

