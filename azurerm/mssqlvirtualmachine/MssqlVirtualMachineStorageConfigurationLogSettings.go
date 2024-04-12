// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine


type MssqlVirtualMachineStorageConfigurationLogSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/mssql_virtual_machine#default_file_path MssqlVirtualMachine#default_file_path}.
	DefaultFilePath *string `field:"required" json:"defaultFilePath" yaml:"defaultFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/mssql_virtual_machine#luns MssqlVirtualMachine#luns}.
	Luns *[]*float64 `field:"required" json:"luns" yaml:"luns"`
}

