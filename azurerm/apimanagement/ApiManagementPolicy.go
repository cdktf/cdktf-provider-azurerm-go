package apimanagement


type ApiManagementPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#xml_content ApiManagement#xml_content}.
	XmlContent *string `field:"optional" json:"xmlContent" yaml:"xmlContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#xml_link ApiManagement#xml_link}.
	XmlLink *string `field:"optional" json:"xmlLink" yaml:"xmlLink"`
}

