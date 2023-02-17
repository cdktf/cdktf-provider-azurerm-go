package mssqlvirtualmachine


type MssqlVirtualMachineStorageConfigurationLogSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#default_file_path MssqlVirtualMachine#default_file_path}.
	DefaultFilePath *string `field:"required" json:"defaultFilePath" yaml:"defaultFilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#luns MssqlVirtualMachine#luns}.
	Luns *[]*float64 `field:"required" json:"luns" yaml:"luns"`
}

