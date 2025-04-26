// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine


type MssqlVirtualMachineStorageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/mssql_virtual_machine#disk_type MssqlVirtualMachine#disk_type}.
	DiskType *string `field:"required" json:"diskType" yaml:"diskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/mssql_virtual_machine#storage_workload_type MssqlVirtualMachine#storage_workload_type}.
	StorageWorkloadType *string `field:"required" json:"storageWorkloadType" yaml:"storageWorkloadType"`
	// data_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/mssql_virtual_machine#data_settings MssqlVirtualMachine#data_settings}
	DataSettings *MssqlVirtualMachineStorageConfigurationDataSettings `field:"optional" json:"dataSettings" yaml:"dataSettings"`
	// log_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/mssql_virtual_machine#log_settings MssqlVirtualMachine#log_settings}
	LogSettings *MssqlVirtualMachineStorageConfigurationLogSettings `field:"optional" json:"logSettings" yaml:"logSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/mssql_virtual_machine#system_db_on_data_disk_enabled MssqlVirtualMachine#system_db_on_data_disk_enabled}.
	SystemDbOnDataDiskEnabled interface{} `field:"optional" json:"systemDbOnDataDiskEnabled" yaml:"systemDbOnDataDiskEnabled"`
	// temp_db_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/mssql_virtual_machine#temp_db_settings MssqlVirtualMachine#temp_db_settings}
	TempDbSettings *MssqlVirtualMachineStorageConfigurationTempDbSettings `field:"optional" json:"tempDbSettings" yaml:"tempDbSettings"`
}

