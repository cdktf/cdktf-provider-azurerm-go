// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupinstancemysqlflexibleserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupInstanceMysqlFlexibleServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/data_protection_backup_instance_mysql_flexible_server#backup_policy_id DataProtectionBackupInstanceMysqlFlexibleServer#backup_policy_id}.
	BackupPolicyId *string `field:"required" json:"backupPolicyId" yaml:"backupPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/data_protection_backup_instance_mysql_flexible_server#location DataProtectionBackupInstanceMysqlFlexibleServer#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/data_protection_backup_instance_mysql_flexible_server#name DataProtectionBackupInstanceMysqlFlexibleServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/data_protection_backup_instance_mysql_flexible_server#server_id DataProtectionBackupInstanceMysqlFlexibleServer#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/data_protection_backup_instance_mysql_flexible_server#vault_id DataProtectionBackupInstanceMysqlFlexibleServer#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/data_protection_backup_instance_mysql_flexible_server#id DataProtectionBackupInstanceMysqlFlexibleServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/data_protection_backup_instance_mysql_flexible_server#timeouts DataProtectionBackupInstanceMysqlFlexibleServer#timeouts}
	Timeouts *DataProtectionBackupInstanceMysqlFlexibleServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

