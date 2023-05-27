package apimanagementdiagnostic


type ApiManagementDiagnosticBackendRequestDataMaskingQueryParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/api_management_diagnostic#mode ApiManagementDiagnostic#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/api_management_diagnostic#value ApiManagementDiagnostic#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

