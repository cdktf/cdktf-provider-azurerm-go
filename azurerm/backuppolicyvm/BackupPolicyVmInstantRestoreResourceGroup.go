package backuppolicyvm


type BackupPolicyVmInstantRestoreResourceGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm#prefix BackupPolicyVm#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm#suffix BackupPolicyVm#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}
