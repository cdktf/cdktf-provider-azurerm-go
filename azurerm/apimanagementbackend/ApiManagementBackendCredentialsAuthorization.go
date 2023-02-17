package apimanagementbackend


type ApiManagementBackendCredentialsAuthorization struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_backend#parameter ApiManagementBackend#parameter}.
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_backend#scheme ApiManagementBackend#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

