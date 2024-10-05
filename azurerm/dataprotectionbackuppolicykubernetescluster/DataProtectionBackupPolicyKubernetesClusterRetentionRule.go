// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicykubernetescluster


type DataProtectionBackupPolicyKubernetesClusterRetentionRule struct {
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#criteria DataProtectionBackupPolicyKubernetesCluster#criteria}
	Criteria *DataProtectionBackupPolicyKubernetesClusterRetentionRuleCriteria `field:"required" json:"criteria" yaml:"criteria"`
	// life_cycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#life_cycle DataProtectionBackupPolicyKubernetesCluster#life_cycle}
	LifeCycle interface{} `field:"required" json:"lifeCycle" yaml:"lifeCycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#name DataProtectionBackupPolicyKubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#priority DataProtectionBackupPolicyKubernetesCluster#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

