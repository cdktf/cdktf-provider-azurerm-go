package servicefabricmanagedcluster


type ServiceFabricManagedClusterAuthenticationCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/service_fabric_managed_cluster#thumbprint ServiceFabricManagedCluster#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/service_fabric_managed_cluster#type ServiceFabricManagedCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/service_fabric_managed_cluster#common_name ServiceFabricManagedCluster#common_name}.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
}

