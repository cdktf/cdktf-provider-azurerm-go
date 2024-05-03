// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightsparkcluster


type HdinsightSparkClusterStorageAccountGen2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/hdinsight_spark_cluster#filesystem_id HdinsightSparkCluster#filesystem_id}.
	FilesystemId *string `field:"required" json:"filesystemId" yaml:"filesystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/hdinsight_spark_cluster#is_default HdinsightSparkCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/hdinsight_spark_cluster#managed_identity_resource_id HdinsightSparkCluster#managed_identity_resource_id}.
	ManagedIdentityResourceId *string `field:"required" json:"managedIdentityResourceId" yaml:"managedIdentityResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/hdinsight_spark_cluster#storage_resource_id HdinsightSparkCluster#storage_resource_id}.
	StorageResourceId *string `field:"required" json:"storageResourceId" yaml:"storageResourceId"`
}

