// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicypostgresql


type DataProtectionBackupPolicyPostgresqlRetentionRule struct {
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_protection_backup_policy_postgresql#criteria DataProtectionBackupPolicyPostgresql#criteria}
	Criteria *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_protection_backup_policy_postgresql#duration DataProtectionBackupPolicyPostgresql#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_protection_backup_policy_postgresql#name DataProtectionBackupPolicyPostgresql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_protection_backup_policy_postgresql#priority DataProtectionBackupPolicyPostgresql#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

