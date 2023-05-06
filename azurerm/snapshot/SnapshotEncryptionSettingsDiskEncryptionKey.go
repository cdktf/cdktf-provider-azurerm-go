package snapshot


type SnapshotEncryptionSettingsDiskEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/snapshot#secret_url Snapshot#secret_url}.
	SecretUrl *string `field:"required" json:"secretUrl" yaml:"secretUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/snapshot#source_vault_id Snapshot#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

