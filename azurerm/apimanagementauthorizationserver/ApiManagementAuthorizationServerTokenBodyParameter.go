package apimanagementauthorizationserver


type ApiManagementAuthorizationServerTokenBodyParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_authorization_server#name ApiManagementAuthorizationServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_authorization_server#value ApiManagementAuthorizationServer#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}
