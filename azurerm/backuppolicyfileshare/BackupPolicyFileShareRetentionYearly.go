package backuppolicyfileshare


type BackupPolicyFileShareRetentionYearly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/backup_policy_file_share#count BackupPolicyFileShare#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/backup_policy_file_share#months BackupPolicyFileShare#months}.
	Months *[]*string `field:"required" json:"months" yaml:"months"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/backup_policy_file_share#weekdays BackupPolicyFileShare#weekdays}.
	Weekdays *[]*string `field:"required" json:"weekdays" yaml:"weekdays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/backup_policy_file_share#weeks BackupPolicyFileShare#weeks}.
	Weeks *[]*string `field:"required" json:"weeks" yaml:"weeks"`
}

