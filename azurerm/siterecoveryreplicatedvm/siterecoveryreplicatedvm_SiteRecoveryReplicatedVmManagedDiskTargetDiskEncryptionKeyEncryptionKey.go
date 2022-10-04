package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#key_url SiteRecoveryReplicatedVm#key_url}.
	KeyUrl *string `field:"optional" json:"keyUrl" yaml:"keyUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#vault_id SiteRecoveryReplicatedVm#vault_id}.
	VaultId *string `field:"optional" json:"vaultId" yaml:"vaultId"`
}

