package servicefabricmanagedcluster


type ServiceFabricManagedClusterNodeTypeVmSecrets struct {
	// certificates block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#certificates ServiceFabricManagedCluster#certificates}
	Certificates interface{} `field:"required" json:"certificates" yaml:"certificates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vault_id ServiceFabricManagedCluster#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
}

