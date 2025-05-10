// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqljobstep


type MssqlJobStepOutputTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_job_step#job_credential_id MssqlJobStep#job_credential_id}.
	JobCredentialId *string `field:"required" json:"jobCredentialId" yaml:"jobCredentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_job_step#mssql_database_id MssqlJobStep#mssql_database_id}.
	MssqlDatabaseId *string `field:"required" json:"mssqlDatabaseId" yaml:"mssqlDatabaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_job_step#table_name MssqlJobStep#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/mssql_job_step#schema_name MssqlJobStep#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

