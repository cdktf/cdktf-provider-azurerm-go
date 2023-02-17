package apimanagementdiagnostic


type ApiManagementDiagnosticBackendResponseDataMaskingHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_diagnostic#mode ApiManagementDiagnostic#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_diagnostic#value ApiManagementDiagnostic#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

