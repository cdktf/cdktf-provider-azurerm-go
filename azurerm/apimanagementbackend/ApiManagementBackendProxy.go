package apimanagementbackend


type ApiManagementBackendProxy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_backend#url ApiManagementBackend#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_backend#username ApiManagementBackend#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_backend#password ApiManagementBackend#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
}

