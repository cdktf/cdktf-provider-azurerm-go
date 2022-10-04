package mssqlvirtualmachine


type MssqlVirtualMachineStorageConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#disk_type MssqlVirtualMachine#disk_type}.
	DiskType *string `field:"required" json:"diskType" yaml:"diskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#storage_workload_type MssqlVirtualMachine#storage_workload_type}.
	StorageWorkloadType *string `field:"required" json:"storageWorkloadType" yaml:"storageWorkloadType"`
	// data_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#data_settings MssqlVirtualMachine#data_settings}
	DataSettings *MssqlVirtualMachineStorageConfigurationDataSettings `field:"optional" json:"dataSettings" yaml:"dataSettings"`
	// log_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#log_settings MssqlVirtualMachine#log_settings}
	LogSettings *MssqlVirtualMachineStorageConfigurationLogSettings `field:"optional" json:"logSettings" yaml:"logSettings"`
	// temp_db_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#temp_db_settings MssqlVirtualMachine#temp_db_settings}
	TempDbSettings *MssqlVirtualMachineStorageConfigurationTempDbSettings `field:"optional" json:"tempDbSettings" yaml:"tempDbSettings"`
}

