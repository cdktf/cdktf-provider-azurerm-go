package snapshot


type SnapshotEncryptionSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/snapshot#enabled Snapshot#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/snapshot#disk_encryption_key Snapshot#disk_encryption_key}
	DiskEncryptionKey *SnapshotEncryptionSettingsDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// key_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/snapshot#key_encryption_key Snapshot#key_encryption_key}
	KeyEncryptionKey *SnapshotEncryptionSettingsKeyEncryptionKey `field:"optional" json:"keyEncryptionKey" yaml:"keyEncryptionKey"`
}

