package servicefabriccluster


type ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/service_fabric_cluster#certificate_common_name ServiceFabricCluster#certificate_common_name}.
	CertificateCommonName *string `field:"required" json:"certificateCommonName" yaml:"certificateCommonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/service_fabric_cluster#certificate_issuer_thumbprint ServiceFabricCluster#certificate_issuer_thumbprint}.
	CertificateIssuerThumbprint *string `field:"optional" json:"certificateIssuerThumbprint" yaml:"certificateIssuerThumbprint"`
}

