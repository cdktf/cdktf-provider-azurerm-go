package backuppolicyvm


type BackupPolicyVmRetentionWeekly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm#count BackupPolicyVm#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm#weekdays BackupPolicyVm#weekdays}.
	Weekdays *[]*string `field:"required" json:"weekdays" yaml:"weekdays"`
}
