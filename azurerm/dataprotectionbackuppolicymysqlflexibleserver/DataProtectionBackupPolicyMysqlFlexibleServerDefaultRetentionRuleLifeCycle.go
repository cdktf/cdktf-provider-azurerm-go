// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicymysqlflexibleserver


type DataProtectionBackupPolicyMysqlFlexibleServerDefaultRetentionRuleLifeCycle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/data_protection_backup_policy_mysql_flexible_server#data_store_type DataProtectionBackupPolicyMysqlFlexibleServer#data_store_type}.
	DataStoreType *string `field:"required" json:"dataStoreType" yaml:"dataStoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/data_protection_backup_policy_mysql_flexible_server#duration DataProtectionBackupPolicyMysqlFlexibleServer#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
}

