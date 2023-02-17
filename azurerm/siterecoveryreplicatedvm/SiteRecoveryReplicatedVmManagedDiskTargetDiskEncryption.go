package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#disk_encryption_key SiteRecoveryReplicatedVm#disk_encryption_key}.
	DiskEncryptionKey interface{} `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#key_encryption_key SiteRecoveryReplicatedVm#key_encryption_key}.
	KeyEncryptionKey interface{} `field:"optional" json:"keyEncryptionKey" yaml:"keyEncryptionKey"`
}

