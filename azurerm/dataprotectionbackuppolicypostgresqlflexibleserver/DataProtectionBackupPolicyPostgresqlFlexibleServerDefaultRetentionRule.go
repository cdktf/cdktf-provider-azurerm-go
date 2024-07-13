// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicypostgresqlflexibleserver


type DataProtectionBackupPolicyPostgresqlFlexibleServerDefaultRetentionRule struct {
	// life_cycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_protection_backup_policy_postgresql_flexible_server#life_cycle DataProtectionBackupPolicyPostgresqlFlexibleServer#life_cycle}
	LifeCycle interface{} `field:"required" json:"lifeCycle" yaml:"lifeCycle"`
}

