// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupvault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupVaultConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#datastore_type DataProtectionBackupVault#datastore_type}.
	DatastoreType *string `field:"required" json:"datastoreType" yaml:"datastoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#location DataProtectionBackupVault#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#name DataProtectionBackupVault#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#redundancy DataProtectionBackupVault#redundancy}.
	Redundancy *string `field:"required" json:"redundancy" yaml:"redundancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#resource_group_name DataProtectionBackupVault#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#cross_region_restore_enabled DataProtectionBackupVault#cross_region_restore_enabled}.
	CrossRegionRestoreEnabled interface{} `field:"optional" json:"crossRegionRestoreEnabled" yaml:"crossRegionRestoreEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#id DataProtectionBackupVault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#identity DataProtectionBackupVault#identity}
	Identity *DataProtectionBackupVaultIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#retention_duration_in_days DataProtectionBackupVault#retention_duration_in_days}.
	RetentionDurationInDays *float64 `field:"optional" json:"retentionDurationInDays" yaml:"retentionDurationInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#soft_delete DataProtectionBackupVault#soft_delete}.
	SoftDelete *string `field:"optional" json:"softDelete" yaml:"softDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#tags DataProtectionBackupVault#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_protection_backup_vault#timeouts DataProtectionBackupVault#timeouts}
	Timeouts *DataProtectionBackupVaultTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

