package servicefabriccluster


type ServiceFabricClusterClientCertificateCommonName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#common_name ServiceFabricCluster#common_name}.
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#is_admin ServiceFabricCluster#is_admin}.
	IsAdmin interface{} `field:"required" json:"isAdmin" yaml:"isAdmin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#issuer_thumbprint ServiceFabricCluster#issuer_thumbprint}.
	IssuerThumbprint *string `field:"optional" json:"issuerThumbprint" yaml:"issuerThumbprint"`
}

