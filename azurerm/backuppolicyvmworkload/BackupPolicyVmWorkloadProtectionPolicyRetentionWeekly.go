package backuppolicyvmworkload


type BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#count BackupPolicyVmWorkload#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#weekdays BackupPolicyVmWorkload#weekdays}.
	Weekdays *[]*string `field:"required" json:"weekdays" yaml:"weekdays"`
}

