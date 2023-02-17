package appservice


type AppServiceLogsApplicationLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#level AppService#level}.
	Level *string `field:"required" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#retention_in_days AppService#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#sas_url AppService#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
}

