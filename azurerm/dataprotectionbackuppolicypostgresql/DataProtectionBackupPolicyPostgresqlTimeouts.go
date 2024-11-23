// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicypostgresql


type DataProtectionBackupPolicyPostgresqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_protection_backup_policy_postgresql#create DataProtectionBackupPolicyPostgresql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_protection_backup_policy_postgresql#delete DataProtectionBackupPolicyPostgresql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_protection_backup_policy_postgresql#read DataProtectionBackupPolicyPostgresql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

