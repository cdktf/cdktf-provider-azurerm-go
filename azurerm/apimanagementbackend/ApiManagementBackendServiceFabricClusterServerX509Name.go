package apimanagementbackend


type ApiManagementBackendServiceFabricClusterServerX509Name struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_backend#issuer_certificate_thumbprint ApiManagementBackend#issuer_certificate_thumbprint}.
	IssuerCertificateThumbprint *string `field:"required" json:"issuerCertificateThumbprint" yaml:"issuerCertificateThumbprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_backend#name ApiManagementBackend#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}
