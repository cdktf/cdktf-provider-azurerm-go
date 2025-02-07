// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/hdinsight_hadoop_cluster#is_default HdinsightHadoopCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/hdinsight_hadoop_cluster#storage_account_key HdinsightHadoopCluster#storage_account_key}.
	StorageAccountKey *string `field:"required" json:"storageAccountKey" yaml:"storageAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/hdinsight_hadoop_cluster#storage_container_id HdinsightHadoopCluster#storage_container_id}.
	StorageContainerId *string `field:"required" json:"storageContainerId" yaml:"storageContainerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/hdinsight_hadoop_cluster#storage_resource_id HdinsightHadoopCluster#storage_resource_id}.
	StorageResourceId *string `field:"optional" json:"storageResourceId" yaml:"storageResourceId"`
}

