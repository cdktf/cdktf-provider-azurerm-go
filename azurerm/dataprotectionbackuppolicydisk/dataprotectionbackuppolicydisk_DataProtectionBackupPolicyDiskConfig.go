package dataprotectionbackuppolicydisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupPolicyDiskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#backup_repeating_time_intervals DataProtectionBackupPolicyDisk#backup_repeating_time_intervals}.
	BackupRepeatingTimeIntervals *[]*string `field:"required" json:"backupRepeatingTimeIntervals" yaml:"backupRepeatingTimeIntervals"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#default_retention_duration DataProtectionBackupPolicyDisk#default_retention_duration}.
	DefaultRetentionDuration *string `field:"required" json:"defaultRetentionDuration" yaml:"defaultRetentionDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#name DataProtectionBackupPolicyDisk#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#vault_id DataProtectionBackupPolicyDisk#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#id DataProtectionBackupPolicyDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#retention_rule DataProtectionBackupPolicyDisk#retention_rule}
	RetentionRule interface{} `field:"optional" json:"retentionRule" yaml:"retentionRule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_disk#timeouts DataProtectionBackupPolicyDisk#timeouts}
	Timeouts *DataProtectionBackupPolicyDiskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
