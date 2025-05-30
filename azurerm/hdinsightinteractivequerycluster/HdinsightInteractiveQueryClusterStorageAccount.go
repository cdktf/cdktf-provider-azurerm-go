// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#is_default HdinsightInteractiveQueryCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#storage_account_key HdinsightInteractiveQueryCluster#storage_account_key}.
	StorageAccountKey *string `field:"required" json:"storageAccountKey" yaml:"storageAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#storage_container_id HdinsightInteractiveQueryCluster#storage_container_id}.
	StorageContainerId *string `field:"required" json:"storageContainerId" yaml:"storageContainerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#storage_resource_id HdinsightInteractiveQueryCluster#storage_resource_id}.
	StorageResourceId *string `field:"optional" json:"storageResourceId" yaml:"storageResourceId"`
}

