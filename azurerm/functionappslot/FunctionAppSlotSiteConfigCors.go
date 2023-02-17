package functionappslot


type FunctionAppSlotSiteConfigCors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/function_app_slot#allowed_origins FunctionAppSlot#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/function_app_slot#support_credentials FunctionAppSlot#support_credentials}.
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

