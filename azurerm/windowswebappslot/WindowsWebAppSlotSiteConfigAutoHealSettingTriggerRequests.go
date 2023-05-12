package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/windows_web_app_slot#count WindowsWebAppSlot#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/windows_web_app_slot#interval WindowsWebAppSlot#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
}

