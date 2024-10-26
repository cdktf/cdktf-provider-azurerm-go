// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicypostgresqlflexibleserver


type DataProtectionBackupPolicyPostgresqlFlexibleServerRetentionRuleLifeCycle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#data_store_type DataProtectionBackupPolicyPostgresqlFlexibleServer#data_store_type}.
	DataStoreType *string `field:"required" json:"dataStoreType" yaml:"dataStoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#duration DataProtectionBackupPolicyPostgresqlFlexibleServer#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
}

