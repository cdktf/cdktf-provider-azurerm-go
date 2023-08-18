package servicefabriccluster


type ServiceFabricClusterCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/service_fabric_cluster#thumbprint ServiceFabricCluster#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/service_fabric_cluster#x509_store_name ServiceFabricCluster#x509_store_name}.
	X509StoreName *string `field:"required" json:"x509StoreName" yaml:"x509StoreName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/service_fabric_cluster#thumbprint_secondary ServiceFabricCluster#thumbprint_secondary}.
	ThumbprintSecondary *string `field:"optional" json:"thumbprintSecondary" yaml:"thumbprintSecondary"`
}

