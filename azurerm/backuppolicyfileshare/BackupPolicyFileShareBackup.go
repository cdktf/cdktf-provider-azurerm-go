package backuppolicyfileshare


type BackupPolicyFileShareBackup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/backup_policy_file_share#frequency BackupPolicyFileShare#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/backup_policy_file_share#time BackupPolicyFileShare#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

