package mssqlvirtualmachine


type MssqlVirtualMachineStorageConfigurationTempDbSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_virtual_machine#default_file_path MssqlVirtualMachine#default_file_path}.
	DefaultFilePath *string `field:"required" json:"defaultFilePath" yaml:"defaultFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_virtual_machine#luns MssqlVirtualMachine#luns}.
	Luns *[]*float64 `field:"required" json:"luns" yaml:"luns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_virtual_machine#data_file_count MssqlVirtualMachine#data_file_count}.
	DataFileCount *float64 `field:"optional" json:"dataFileCount" yaml:"dataFileCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_virtual_machine#data_file_growth_in_mb MssqlVirtualMachine#data_file_growth_in_mb}.
	DataFileGrowthInMb *float64 `field:"optional" json:"dataFileGrowthInMb" yaml:"dataFileGrowthInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_virtual_machine#data_file_size_mb MssqlVirtualMachine#data_file_size_mb}.
	DataFileSizeMb *float64 `field:"optional" json:"dataFileSizeMb" yaml:"dataFileSizeMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_virtual_machine#log_file_growth_mb MssqlVirtualMachine#log_file_growth_mb}.
	LogFileGrowthMb *float64 `field:"optional" json:"logFileGrowthMb" yaml:"logFileGrowthMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_virtual_machine#log_file_size_mb MssqlVirtualMachine#log_file_size_mb}.
	LogFileSizeMb *float64 `field:"optional" json:"logFileSizeMb" yaml:"logFileSizeMb"`
}

