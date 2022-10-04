package backupprotectedfileshare


type BackupProtectedFileShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_protected_file_share#create BackupProtectedFileShare#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_protected_file_share#delete BackupProtectedFileShare#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_protected_file_share#read BackupProtectedFileShare#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_protected_file_share#update BackupProtectedFileShare#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

