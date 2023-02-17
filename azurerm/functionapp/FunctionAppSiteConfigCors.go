package functionapp


type FunctionAppSiteConfigCors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/function_app#allowed_origins FunctionApp#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/function_app#support_credentials FunctionApp#support_credentials}.
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

