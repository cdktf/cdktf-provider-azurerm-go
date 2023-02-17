package dataazurermlogicappstandard


type DataAzurermLogicAppStandardSiteConfigCors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/logic_app_standard#allowed_origins DataAzurermLogicAppStandard#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/logic_app_standard#support_credentials DataAzurermLogicAppStandard#support_credentials}.
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

