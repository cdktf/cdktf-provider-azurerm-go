package backuppolicyvm


type BackupPolicyVmRetentionYearly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm#count BackupPolicyVm#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm#months BackupPolicyVm#months}.
	Months *[]*string `field:"required" json:"months" yaml:"months"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm#weekdays BackupPolicyVm#weekdays}.
	Weekdays *[]*string `field:"required" json:"weekdays" yaml:"weekdays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm#weeks BackupPolicyVm#weeks}.
	Weeks *[]*string `field:"required" json:"weeks" yaml:"weeks"`
}
