package snapshot


type SnapshotEncryptionSettingsKeyEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/snapshot#key_url Snapshot#key_url}.
	KeyUrl *string `field:"required" json:"keyUrl" yaml:"keyUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/snapshot#source_vault_id Snapshot#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

