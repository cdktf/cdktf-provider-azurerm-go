package dataprotectionbackuppolicypostgresql


type DataProtectionBackupPolicyPostgresqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#create DataProtectionBackupPolicyPostgresql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#delete DataProtectionBackupPolicyPostgresql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#read DataProtectionBackupPolicyPostgresql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#update DataProtectionBackupPolicyPostgresql#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
