package servicefabriccluster


type ServiceFabricClusterClientCertificateThumbprint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#is_admin ServiceFabricCluster#is_admin}.
	IsAdmin interface{} `field:"required" json:"isAdmin" yaml:"isAdmin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#thumbprint ServiceFabricCluster#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
}

