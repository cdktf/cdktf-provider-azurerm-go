package backuppolicyvmworkload


type BackupPolicyVmWorkloadProtectionPolicy struct {
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/backup_policy_vm_workload#backup BackupPolicyVmWorkload#backup}
	Backup *BackupPolicyVmWorkloadProtectionPolicyBackup `field:"required" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/backup_policy_vm_workload#policy_type BackupPolicyVmWorkload#policy_type}.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// retention_daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/backup_policy_vm_workload#retention_daily BackupPolicyVmWorkload#retention_daily}
	RetentionDaily *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily `field:"optional" json:"retentionDaily" yaml:"retentionDaily"`
	// retention_monthly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/backup_policy_vm_workload#retention_monthly BackupPolicyVmWorkload#retention_monthly}
	RetentionMonthly *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly `field:"optional" json:"retentionMonthly" yaml:"retentionMonthly"`
	// retention_weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/backup_policy_vm_workload#retention_weekly BackupPolicyVmWorkload#retention_weekly}
	RetentionWeekly *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly `field:"optional" json:"retentionWeekly" yaml:"retentionWeekly"`
	// retention_yearly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/backup_policy_vm_workload#retention_yearly BackupPolicyVmWorkload#retention_yearly}
	RetentionYearly *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly `field:"optional" json:"retentionYearly" yaml:"retentionYearly"`
	// simple_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/backup_policy_vm_workload#simple_retention BackupPolicyVmWorkload#simple_retention}
	SimpleRetention *BackupPolicyVmWorkloadProtectionPolicySimpleRetention `field:"optional" json:"simpleRetention" yaml:"simpleRetention"`
}

