package mssqlvirtualmachine


type MssqlVirtualMachineSqlInstance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#adhoc_workloads_optimization_enabled MssqlVirtualMachine#adhoc_workloads_optimization_enabled}.
	AdhocWorkloadsOptimizationEnabled interface{} `field:"optional" json:"adhocWorkloadsOptimizationEnabled" yaml:"adhocWorkloadsOptimizationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#collation MssqlVirtualMachine#collation}.
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#instant_file_initialization_enabled MssqlVirtualMachine#instant_file_initialization_enabled}.
	InstantFileInitializationEnabled interface{} `field:"optional" json:"instantFileInitializationEnabled" yaml:"instantFileInitializationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#lock_pages_in_memory_enabled MssqlVirtualMachine#lock_pages_in_memory_enabled}.
	LockPagesInMemoryEnabled interface{} `field:"optional" json:"lockPagesInMemoryEnabled" yaml:"lockPagesInMemoryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#max_dop MssqlVirtualMachine#max_dop}.
	MaxDop *float64 `field:"optional" json:"maxDop" yaml:"maxDop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#max_server_memory_mb MssqlVirtualMachine#max_server_memory_mb}.
	MaxServerMemoryMb *float64 `field:"optional" json:"maxServerMemoryMb" yaml:"maxServerMemoryMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#min_server_memory_mb MssqlVirtualMachine#min_server_memory_mb}.
	MinServerMemoryMb *float64 `field:"optional" json:"minServerMemoryMb" yaml:"minServerMemoryMb"`
}

