package linuxwebappslot


type LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_web_app_slot#count LinuxWebAppSlot#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_web_app_slot#interval LinuxWebAppSlot#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
}

