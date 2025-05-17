// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqldatabase


type MssqlDatabaseImport struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_database#administrator_login MssqlDatabase#administrator_login}.
	AdministratorLogin *string `field:"required" json:"administratorLogin" yaml:"administratorLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_database#administrator_login_password MssqlDatabase#administrator_login_password}.
	AdministratorLoginPassword *string `field:"required" json:"administratorLoginPassword" yaml:"administratorLoginPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_database#authentication_type MssqlDatabase#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_database#storage_key MssqlDatabase#storage_key}.
	StorageKey *string `field:"required" json:"storageKey" yaml:"storageKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_database#storage_key_type MssqlDatabase#storage_key_type}.
	StorageKeyType *string `field:"required" json:"storageKeyType" yaml:"storageKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_database#storage_uri MssqlDatabase#storage_uri}.
	StorageUri *string `field:"required" json:"storageUri" yaml:"storageUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_database#storage_account_id MssqlDatabase#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
}

