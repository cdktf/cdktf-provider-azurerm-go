package apimanagementapi


type ApiManagementApiImport struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api#content_format ApiManagementApi#content_format}.
	ContentFormat *string `field:"required" json:"contentFormat" yaml:"contentFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api#content_value ApiManagementApi#content_value}.
	ContentValue *string `field:"required" json:"contentValue" yaml:"contentValue"`
	// wsdl_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api#wsdl_selector ApiManagementApi#wsdl_selector}
	WsdlSelector *ApiManagementApiImportWsdlSelector `field:"optional" json:"wsdlSelector" yaml:"wsdlSelector"`
}
