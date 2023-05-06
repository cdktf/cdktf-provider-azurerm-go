package appserviceslot


type AppServiceSlotLogsHttpLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/app_service_slot#retention_in_days AppServiceSlot#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/app_service_slot#sas_url AppServiceSlot#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
}

