package linuxwebappslot


type LinuxWebAppSlotSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#action LinuxWebAppSlot#action}
	Action *LinuxWebAppSlotSiteConfigAutoHealSettingAction `field:"optional" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#trigger LinuxWebAppSlot#trigger}
	Trigger *LinuxWebAppSlotSiteConfigAutoHealSettingTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}
