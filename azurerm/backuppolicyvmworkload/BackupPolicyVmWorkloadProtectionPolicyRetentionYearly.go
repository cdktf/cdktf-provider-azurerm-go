package backuppolicyvmworkload


type BackupPolicyVmWorkloadProtectionPolicyRetentionYearly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/backup_policy_vm_workload#count BackupPolicyVmWorkload#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/backup_policy_vm_workload#format_type BackupPolicyVmWorkload#format_type}.
	FormatType *string `field:"required" json:"formatType" yaml:"formatType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/backup_policy_vm_workload#months BackupPolicyVmWorkload#months}.
	Months *[]*string `field:"required" json:"months" yaml:"months"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/backup_policy_vm_workload#monthdays BackupPolicyVmWorkload#monthdays}.
	Monthdays *[]*float64 `field:"optional" json:"monthdays" yaml:"monthdays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/backup_policy_vm_workload#weekdays BackupPolicyVmWorkload#weekdays}.
	Weekdays *[]*string `field:"optional" json:"weekdays" yaml:"weekdays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/backup_policy_vm_workload#weeks BackupPolicyVmWorkload#weeks}.
	Weeks *[]*string `field:"optional" json:"weeks" yaml:"weeks"`
}

