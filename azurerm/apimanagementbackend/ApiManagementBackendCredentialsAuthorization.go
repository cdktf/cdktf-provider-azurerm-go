package apimanagementbackend


type ApiManagementBackendCredentialsAuthorization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/api_management_backend#parameter ApiManagementBackend#parameter}.
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/api_management_backend#scheme ApiManagementBackend#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

