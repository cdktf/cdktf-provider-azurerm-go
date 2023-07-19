package keyvaultaccesspolicy


type KeyVaultAccessPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/key_vault_access_policy#create KeyVaultAccessPolicyA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/key_vault_access_policy#delete KeyVaultAccessPolicyA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/key_vault_access_policy#read KeyVaultAccessPolicyA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/key_vault_access_policy#update KeyVaultAccessPolicyA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

