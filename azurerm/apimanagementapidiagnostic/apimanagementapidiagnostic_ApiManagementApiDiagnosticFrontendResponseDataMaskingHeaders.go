package apimanagementapidiagnostic


type ApiManagementApiDiagnosticFrontendResponseDataMaskingHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api_diagnostic#mode ApiManagementApiDiagnostic#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api_diagnostic#value ApiManagementApiDiagnostic#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}
