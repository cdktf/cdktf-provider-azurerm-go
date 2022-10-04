package dataprotectionbackuppolicydisk


type DataProtectionBackupPolicyDiskTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#create DataProtectionBackupPolicyDisk#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#delete DataProtectionBackupPolicyDisk#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#read DataProtectionBackupPolicyDisk#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#update DataProtectionBackupPolicyDisk#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

