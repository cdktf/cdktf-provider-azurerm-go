package windowswebappslot


type WindowsWebAppSlotLogsHttpLogsFileSystem struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#retention_in_days WindowsWebAppSlot#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#retention_in_mb WindowsWebAppSlot#retention_in_mb}.
	RetentionInMb *float64 `field:"required" json:"retentionInMb" yaml:"retentionInMb"`
}

