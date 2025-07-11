// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupinstancepostgresql

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupInstancePostgresqlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_protection_backup_instance_postgresql#backup_policy_id DataProtectionBackupInstancePostgresql#backup_policy_id}.
	BackupPolicyId *string `field:"required" json:"backupPolicyId" yaml:"backupPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_protection_backup_instance_postgresql#database_id DataProtectionBackupInstancePostgresql#database_id}.
	DatabaseId *string `field:"required" json:"databaseId" yaml:"databaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_protection_backup_instance_postgresql#location DataProtectionBackupInstancePostgresql#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_protection_backup_instance_postgresql#name DataProtectionBackupInstancePostgresql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_protection_backup_instance_postgresql#vault_id DataProtectionBackupInstancePostgresql#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_protection_backup_instance_postgresql#database_credential_key_vault_secret_id DataProtectionBackupInstancePostgresql#database_credential_key_vault_secret_id}.
	DatabaseCredentialKeyVaultSecretId *string `field:"optional" json:"databaseCredentialKeyVaultSecretId" yaml:"databaseCredentialKeyVaultSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_protection_backup_instance_postgresql#id DataProtectionBackupInstancePostgresql#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_protection_backup_instance_postgresql#timeouts DataProtectionBackupInstancePostgresql#timeouts}
	Timeouts *DataProtectionBackupInstancePostgresqlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

