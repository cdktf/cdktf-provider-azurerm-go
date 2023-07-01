package linuxfunctionappslot


type LinuxFunctionAppSlotSiteConfigAppServiceLogs struct {
	// The amount of disk space to use for logs. Valid values are between `25` and `100`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_function_app_slot#disk_quota_mb LinuxFunctionAppSlot#disk_quota_mb}
	DiskQuotaMb *float64 `field:"optional" json:"diskQuotaMb" yaml:"diskQuotaMb"`
	// The retention period for logs in days. Valid values are between `0` and `99999`. Defaults to `0` (never delete).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_function_app_slot#retention_period_days LinuxFunctionAppSlot#retention_period_days}
	RetentionPeriodDays *float64 `field:"optional" json:"retentionPeriodDays" yaml:"retentionPeriodDays"`
}

