package backuppolicyvmworkload


type BackupPolicyVmWorkloadTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#create BackupPolicyVmWorkload#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#delete BackupPolicyVmWorkload#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#read BackupPolicyVmWorkload#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#update BackupPolicyVmWorkload#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
