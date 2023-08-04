package batchpool


type BatchPoolDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/batch_pool#disk_encryption_target BatchPool#disk_encryption_target}.
	DiskEncryptionTarget *string `field:"required" json:"diskEncryptionTarget" yaml:"diskEncryptionTarget"`
}

