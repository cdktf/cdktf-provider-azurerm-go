package backuppolicyfileshare


type BackupPolicyFileShareRetentionMonthly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_file_share#count BackupPolicyFileShare#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_file_share#weekdays BackupPolicyFileShare#weekdays}.
	Weekdays *[]*string `field:"required" json:"weekdays" yaml:"weekdays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_file_share#weeks BackupPolicyFileShare#weeks}.
	Weeks *[]*string `field:"required" json:"weeks" yaml:"weeks"`
}

