package apimanagementproducttag


type ApiManagementProductTagTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_product_tag#create ApiManagementProductTag#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_product_tag#delete ApiManagementProductTag#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_product_tag#read ApiManagementProductTag#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_product_tag#update ApiManagementProductTag#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
