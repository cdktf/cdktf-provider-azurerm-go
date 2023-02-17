package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#secret_url SiteRecoveryReplicatedVm#secret_url}.
	SecretUrl *string `field:"optional" json:"secretUrl" yaml:"secretUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#vault_id SiteRecoveryReplicatedVm#vault_id}.
	VaultId *string `field:"optional" json:"vaultId" yaml:"vaultId"`
}

