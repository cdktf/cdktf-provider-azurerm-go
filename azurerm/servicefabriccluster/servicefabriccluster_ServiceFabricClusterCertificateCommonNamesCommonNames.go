package servicefabriccluster


type ServiceFabricClusterCertificateCommonNamesCommonNames struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#certificate_common_name ServiceFabricCluster#certificate_common_name}.
	CertificateCommonName *string `field:"required" json:"certificateCommonName" yaml:"certificateCommonName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#certificate_issuer_thumbprint ServiceFabricCluster#certificate_issuer_thumbprint}.
	CertificateIssuerThumbprint *string `field:"optional" json:"certificateIssuerThumbprint" yaml:"certificateIssuerThumbprint"`
}

