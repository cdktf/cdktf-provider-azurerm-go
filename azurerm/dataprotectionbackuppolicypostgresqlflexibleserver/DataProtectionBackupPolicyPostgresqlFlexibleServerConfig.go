// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicypostgresqlflexibleserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupPolicyPostgresqlFlexibleServerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#backup_repeating_time_intervals DataProtectionBackupPolicyPostgresqlFlexibleServer#backup_repeating_time_intervals}.
	BackupRepeatingTimeIntervals *[]*string `field:"required" json:"backupRepeatingTimeIntervals" yaml:"backupRepeatingTimeIntervals"`
	// default_retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#default_retention_rule DataProtectionBackupPolicyPostgresqlFlexibleServer#default_retention_rule}
	DefaultRetentionRule *DataProtectionBackupPolicyPostgresqlFlexibleServerDefaultRetentionRule `field:"required" json:"defaultRetentionRule" yaml:"defaultRetentionRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#name DataProtectionBackupPolicyPostgresqlFlexibleServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#vault_id DataProtectionBackupPolicyPostgresqlFlexibleServer#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#id DataProtectionBackupPolicyPostgresqlFlexibleServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#retention_rule DataProtectionBackupPolicyPostgresqlFlexibleServer#retention_rule}
	RetentionRule interface{} `field:"optional" json:"retentionRule" yaml:"retentionRule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#timeouts DataProtectionBackupPolicyPostgresqlFlexibleServer#timeouts}
	Timeouts *DataProtectionBackupPolicyPostgresqlFlexibleServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#time_zone DataProtectionBackupPolicyPostgresqlFlexibleServer#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

