package apimanagementglobalschema


type ApiManagementGlobalSchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_global_schema#create ApiManagementGlobalSchema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_global_schema#delete ApiManagementGlobalSchema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_global_schema#read ApiManagementGlobalSchema#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_global_schema#update ApiManagementGlobalSchema#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
