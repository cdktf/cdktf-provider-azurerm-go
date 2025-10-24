// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracleautonomousdatabaseclonefrombackup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleAutonomousDatabaseCloneFromBackupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#admin_password OracleAutonomousDatabaseCloneFromBackup#admin_password}.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#auto_scaling_enabled OracleAutonomousDatabaseCloneFromBackup#auto_scaling_enabled}.
	AutoScalingEnabled interface{} `field:"required" json:"autoScalingEnabled" yaml:"autoScalingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#auto_scaling_for_storage_enabled OracleAutonomousDatabaseCloneFromBackup#auto_scaling_for_storage_enabled}.
	AutoScalingForStorageEnabled interface{} `field:"required" json:"autoScalingForStorageEnabled" yaml:"autoScalingForStorageEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#backup_retention_period_in_days OracleAutonomousDatabaseCloneFromBackup#backup_retention_period_in_days}.
	BackupRetentionPeriodInDays *float64 `field:"required" json:"backupRetentionPeriodInDays" yaml:"backupRetentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#character_set OracleAutonomousDatabaseCloneFromBackup#character_set}.
	CharacterSet *string `field:"required" json:"characterSet" yaml:"characterSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#clone_type OracleAutonomousDatabaseCloneFromBackup#clone_type}.
	CloneType *string `field:"required" json:"cloneType" yaml:"cloneType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#compute_count OracleAutonomousDatabaseCloneFromBackup#compute_count}.
	ComputeCount *float64 `field:"required" json:"computeCount" yaml:"computeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#compute_model OracleAutonomousDatabaseCloneFromBackup#compute_model}.
	ComputeModel *string `field:"required" json:"computeModel" yaml:"computeModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#database_version OracleAutonomousDatabaseCloneFromBackup#database_version}.
	DatabaseVersion *string `field:"required" json:"databaseVersion" yaml:"databaseVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#database_workload OracleAutonomousDatabaseCloneFromBackup#database_workload}.
	DatabaseWorkload *string `field:"required" json:"databaseWorkload" yaml:"databaseWorkload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#data_storage_size_in_tb OracleAutonomousDatabaseCloneFromBackup#data_storage_size_in_tb}.
	DataStorageSizeInTb *float64 `field:"required" json:"dataStorageSizeInTb" yaml:"dataStorageSizeInTb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#display_name OracleAutonomousDatabaseCloneFromBackup#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#license_model OracleAutonomousDatabaseCloneFromBackup#license_model}.
	LicenseModel *string `field:"required" json:"licenseModel" yaml:"licenseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#location OracleAutonomousDatabaseCloneFromBackup#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#mtls_connection_required OracleAutonomousDatabaseCloneFromBackup#mtls_connection_required}.
	MtlsConnectionRequired interface{} `field:"required" json:"mtlsConnectionRequired" yaml:"mtlsConnectionRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#name OracleAutonomousDatabaseCloneFromBackup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#national_character_set OracleAutonomousDatabaseCloneFromBackup#national_character_set}.
	NationalCharacterSet *string `field:"required" json:"nationalCharacterSet" yaml:"nationalCharacterSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#resource_group_name OracleAutonomousDatabaseCloneFromBackup#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#source_autonomous_database_id OracleAutonomousDatabaseCloneFromBackup#source_autonomous_database_id}.
	SourceAutonomousDatabaseId *string `field:"required" json:"sourceAutonomousDatabaseId" yaml:"sourceAutonomousDatabaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#allowed_ip_addresses OracleAutonomousDatabaseCloneFromBackup#allowed_ip_addresses}.
	AllowedIpAddresses *[]*string `field:"optional" json:"allowedIpAddresses" yaml:"allowedIpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#backup_timestamp OracleAutonomousDatabaseCloneFromBackup#backup_timestamp}.
	BackupTimestamp *string `field:"optional" json:"backupTimestamp" yaml:"backupTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#customer_contacts OracleAutonomousDatabaseCloneFromBackup#customer_contacts}.
	CustomerContacts *[]*string `field:"optional" json:"customerContacts" yaml:"customerContacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#id OracleAutonomousDatabaseCloneFromBackup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#subnet_id OracleAutonomousDatabaseCloneFromBackup#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#tags OracleAutonomousDatabaseCloneFromBackup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#timeouts OracleAutonomousDatabaseCloneFromBackup#timeouts}
	Timeouts *OracleAutonomousDatabaseCloneFromBackupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/oracle_autonomous_database_clone_from_backup#virtual_network_id OracleAutonomousDatabaseCloneFromBackup#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

