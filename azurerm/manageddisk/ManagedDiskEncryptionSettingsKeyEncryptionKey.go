package manageddisk


type ManagedDiskEncryptionSettingsKeyEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/managed_disk#key_url ManagedDisk#key_url}.
	KeyUrl *string `field:"required" json:"keyUrl" yaml:"keyUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/managed_disk#source_vault_id ManagedDisk#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

