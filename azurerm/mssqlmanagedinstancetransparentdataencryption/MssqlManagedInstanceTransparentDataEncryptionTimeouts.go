// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstancetransparentdataencryption


type MssqlManagedInstanceTransparentDataEncryptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/mssql_managed_instance_transparent_data_encryption#create MssqlManagedInstanceTransparentDataEncryption#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/mssql_managed_instance_transparent_data_encryption#delete MssqlManagedInstanceTransparentDataEncryption#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/mssql_managed_instance_transparent_data_encryption#read MssqlManagedInstanceTransparentDataEncryption#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/mssql_managed_instance_transparent_data_encryption#update MssqlManagedInstanceTransparentDataEncryption#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

