// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicykubernetescluster


type DataProtectionBackupPolicyKubernetesClusterDefaultRetentionRule struct {
	// life_cycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#life_cycle DataProtectionBackupPolicyKubernetesCluster#life_cycle}
	LifeCycle interface{} `field:"required" json:"lifeCycle" yaml:"lifeCycle"`
}

