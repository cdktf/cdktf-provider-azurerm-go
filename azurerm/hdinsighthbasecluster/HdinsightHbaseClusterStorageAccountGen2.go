package hdinsighthbasecluster


type HdinsightHbaseClusterStorageAccountGen2 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#filesystem_id HdinsightHbaseCluster#filesystem_id}.
	FilesystemId *string `field:"required" json:"filesystemId" yaml:"filesystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#is_default HdinsightHbaseCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#managed_identity_resource_id HdinsightHbaseCluster#managed_identity_resource_id}.
	ManagedIdentityResourceId *string `field:"required" json:"managedIdentityResourceId" yaml:"managedIdentityResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#storage_resource_id HdinsightHbaseCluster#storage_resource_id}.
	StorageResourceId *string `field:"required" json:"storageResourceId" yaml:"storageResourceId"`
}

