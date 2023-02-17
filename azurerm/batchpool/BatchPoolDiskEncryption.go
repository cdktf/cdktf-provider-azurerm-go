package batchpool


type BatchPoolDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#disk_encryption_target BatchPool#disk_encryption_target}.
	DiskEncryptionTarget *string `field:"required" json:"diskEncryptionTarget" yaml:"diskEncryptionTarget"`
}

