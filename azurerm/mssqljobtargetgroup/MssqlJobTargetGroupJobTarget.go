// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqljobtargetgroup


type MssqlJobTargetGroupJobTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/mssql_job_target_group#server_name MssqlJobTargetGroup#server_name}.
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/mssql_job_target_group#database_name MssqlJobTargetGroup#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/mssql_job_target_group#elastic_pool_name MssqlJobTargetGroup#elastic_pool_name}.
	ElasticPoolName *string `field:"optional" json:"elasticPoolName" yaml:"elasticPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/mssql_job_target_group#job_credential_id MssqlJobTargetGroup#job_credential_id}.
	JobCredentialId *string `field:"optional" json:"jobCredentialId" yaml:"jobCredentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/mssql_job_target_group#membership_type MssqlJobTargetGroup#membership_type}.
	MembershipType *string `field:"optional" json:"membershipType" yaml:"membershipType"`
}

