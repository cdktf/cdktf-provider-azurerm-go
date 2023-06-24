package mediaservicesaccount


type MediaServicesAccountEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_services_account#key_vault_key_identifier MediaServicesAccount#key_vault_key_identifier}.
	KeyVaultKeyIdentifier *string `field:"optional" json:"keyVaultKeyIdentifier" yaml:"keyVaultKeyIdentifier"`
	// managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_services_account#managed_identity MediaServicesAccount#managed_identity}
	ManagedIdentity *MediaServicesAccountEncryptionManagedIdentity `field:"optional" json:"managedIdentity" yaml:"managedIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_services_account#type MediaServicesAccount#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

