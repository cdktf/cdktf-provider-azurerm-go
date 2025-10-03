// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracleautonomousdatabasebackup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleAutonomousDatabaseBackupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/oracle_autonomous_database_backup#autonomous_database_id OracleAutonomousDatabaseBackup#autonomous_database_id}.
	AutonomousDatabaseId *string `field:"required" json:"autonomousDatabaseId" yaml:"autonomousDatabaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/oracle_autonomous_database_backup#name OracleAutonomousDatabaseBackup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/oracle_autonomous_database_backup#retention_period_in_days OracleAutonomousDatabaseBackup#retention_period_in_days}.
	RetentionPeriodInDays *float64 `field:"required" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/oracle_autonomous_database_backup#id OracleAutonomousDatabaseBackup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/oracle_autonomous_database_backup#timeouts OracleAutonomousDatabaseBackup#timeouts}
	Timeouts *OracleAutonomousDatabaseBackupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/oracle_autonomous_database_backup#type OracleAutonomousDatabaseBackup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

