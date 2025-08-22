// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicykubernetescluster


type DataProtectionBackupPolicyKubernetesClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#create DataProtectionBackupPolicyKubernetesCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#delete DataProtectionBackupPolicyKubernetesCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#read DataProtectionBackupPolicyKubernetesCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

