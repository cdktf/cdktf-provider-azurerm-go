package manageddisk


type ManagedDiskEncryptionSettings struct {
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#disk_encryption_key ManagedDisk#disk_encryption_key}
	DiskEncryptionKey *ManagedDiskEncryptionSettingsDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#enabled ManagedDisk#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// key_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk#key_encryption_key ManagedDisk#key_encryption_key}
	KeyEncryptionKey *ManagedDiskEncryptionSettingsKeyEncryptionKey `field:"optional" json:"keyEncryptionKey" yaml:"keyEncryptionKey"`
}
