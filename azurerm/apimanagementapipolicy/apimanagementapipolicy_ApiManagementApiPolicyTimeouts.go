package apimanagementapipolicy


type ApiManagementApiPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api_policy#create ApiManagementApiPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api_policy#delete ApiManagementApiPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api_policy#read ApiManagementApiPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api_policy#update ApiManagementApiPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
