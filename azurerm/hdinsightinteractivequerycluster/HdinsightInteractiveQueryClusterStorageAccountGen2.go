package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterStorageAccountGen2 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#filesystem_id HdinsightInteractiveQueryCluster#filesystem_id}.
	FilesystemId *string `field:"required" json:"filesystemId" yaml:"filesystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#is_default HdinsightInteractiveQueryCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#managed_identity_resource_id HdinsightInteractiveQueryCluster#managed_identity_resource_id}.
	ManagedIdentityResourceId *string `field:"required" json:"managedIdentityResourceId" yaml:"managedIdentityResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#storage_resource_id HdinsightInteractiveQueryCluster#storage_resource_id}.
	StorageResourceId *string `field:"required" json:"storageResourceId" yaml:"storageResourceId"`
}

