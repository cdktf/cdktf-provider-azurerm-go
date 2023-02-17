package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#action WindowsWebAppSlot#action}
	Action *WindowsWebAppSlotSiteConfigAutoHealSettingAction `field:"required" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#trigger WindowsWebAppSlot#trigger}
	Trigger *WindowsWebAppSlotSiteConfigAutoHealSettingTrigger `field:"required" json:"trigger" yaml:"trigger"`
}

