package backuppolicyvmworkload


type BackupPolicyVmWorkloadSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/backup_policy_vm_workload#time_zone BackupPolicyVmWorkload#time_zone}.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/backup_policy_vm_workload#compression_enabled BackupPolicyVmWorkload#compression_enabled}.
	CompressionEnabled interface{} `field:"optional" json:"compressionEnabled" yaml:"compressionEnabled"`
}

