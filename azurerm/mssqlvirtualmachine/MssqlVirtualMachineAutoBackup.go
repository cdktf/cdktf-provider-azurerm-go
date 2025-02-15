// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine


type MssqlVirtualMachineAutoBackup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mssql_virtual_machine#retention_period_in_days MssqlVirtualMachine#retention_period_in_days}.
	RetentionPeriodInDays *float64 `field:"required" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mssql_virtual_machine#storage_account_access_key MssqlVirtualMachine#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"required" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mssql_virtual_machine#storage_blob_endpoint MssqlVirtualMachine#storage_blob_endpoint}.
	StorageBlobEndpoint *string `field:"required" json:"storageBlobEndpoint" yaml:"storageBlobEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mssql_virtual_machine#encryption_enabled MssqlVirtualMachine#encryption_enabled}.
	EncryptionEnabled interface{} `field:"optional" json:"encryptionEnabled" yaml:"encryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mssql_virtual_machine#encryption_password MssqlVirtualMachine#encryption_password}.
	EncryptionPassword *string `field:"optional" json:"encryptionPassword" yaml:"encryptionPassword"`
	// manual_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mssql_virtual_machine#manual_schedule MssqlVirtualMachine#manual_schedule}
	ManualSchedule *MssqlVirtualMachineAutoBackupManualSchedule `field:"optional" json:"manualSchedule" yaml:"manualSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/mssql_virtual_machine#system_databases_backup_enabled MssqlVirtualMachine#system_databases_backup_enabled}.
	SystemDatabasesBackupEnabled interface{} `field:"optional" json:"systemDatabasesBackupEnabled" yaml:"systemDatabasesBackupEnabled"`
}

