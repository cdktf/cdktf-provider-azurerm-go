package apimanagementapioperation


type ApiManagementApiOperationRequestRepresentation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/api_management_api_operation#content_type ApiManagementApiOperation#content_type}.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// example block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/api_management_api_operation#example ApiManagementApiOperation#example}
	Example interface{} `field:"optional" json:"example" yaml:"example"`
	// form_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/api_management_api_operation#form_parameter ApiManagementApiOperation#form_parameter}
	FormParameter interface{} `field:"optional" json:"formParameter" yaml:"formParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/api_management_api_operation#schema_id ApiManagementApiOperation#schema_id}.
	SchemaId *string `field:"optional" json:"schemaId" yaml:"schemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/api_management_api_operation#type_name ApiManagementApiOperation#type_name}.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

