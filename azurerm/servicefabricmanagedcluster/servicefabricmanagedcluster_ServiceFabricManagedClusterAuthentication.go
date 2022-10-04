package servicefabricmanagedcluster


type ServiceFabricManagedClusterAuthentication struct {
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#active_directory ServiceFabricManagedCluster#active_directory}
	ActiveDirectory *ServiceFabricManagedClusterAuthenticationActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#certificate ServiceFabricManagedCluster#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
}

