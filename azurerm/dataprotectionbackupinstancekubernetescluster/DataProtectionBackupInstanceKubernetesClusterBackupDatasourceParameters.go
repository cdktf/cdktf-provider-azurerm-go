// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupinstancekubernetescluster


type DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_protection_backup_instance_kubernetes_cluster#cluster_scoped_resources_enabled DataProtectionBackupInstanceKubernetesCluster#cluster_scoped_resources_enabled}.
	ClusterScopedResourcesEnabled interface{} `field:"optional" json:"clusterScopedResourcesEnabled" yaml:"clusterScopedResourcesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_protection_backup_instance_kubernetes_cluster#excluded_namespaces DataProtectionBackupInstanceKubernetesCluster#excluded_namespaces}.
	ExcludedNamespaces *[]*string `field:"optional" json:"excludedNamespaces" yaml:"excludedNamespaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_protection_backup_instance_kubernetes_cluster#excluded_resource_types DataProtectionBackupInstanceKubernetesCluster#excluded_resource_types}.
	ExcludedResourceTypes *[]*string `field:"optional" json:"excludedResourceTypes" yaml:"excludedResourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_protection_backup_instance_kubernetes_cluster#included_namespaces DataProtectionBackupInstanceKubernetesCluster#included_namespaces}.
	IncludedNamespaces *[]*string `field:"optional" json:"includedNamespaces" yaml:"includedNamespaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_protection_backup_instance_kubernetes_cluster#included_resource_types DataProtectionBackupInstanceKubernetesCluster#included_resource_types}.
	IncludedResourceTypes *[]*string `field:"optional" json:"includedResourceTypes" yaml:"includedResourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_protection_backup_instance_kubernetes_cluster#label_selectors DataProtectionBackupInstanceKubernetesCluster#label_selectors}.
	LabelSelectors *[]*string `field:"optional" json:"labelSelectors" yaml:"labelSelectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_protection_backup_instance_kubernetes_cluster#volume_snapshot_enabled DataProtectionBackupInstanceKubernetesCluster#volume_snapshot_enabled}.
	VolumeSnapshotEnabled interface{} `field:"optional" json:"volumeSnapshotEnabled" yaml:"volumeSnapshotEnabled"`
}

