package windowswebappslot


type WindowsWebAppSlotLogsHttpLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#sas_url WindowsWebAppSlot#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#retention_in_days WindowsWebAppSlot#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
}

