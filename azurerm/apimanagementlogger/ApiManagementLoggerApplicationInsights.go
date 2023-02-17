package apimanagementlogger


type ApiManagementLoggerApplicationInsights struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_logger#instrumentation_key ApiManagementLogger#instrumentation_key}.
	InstrumentationKey *string `field:"required" json:"instrumentationKey" yaml:"instrumentationKey"`
}

