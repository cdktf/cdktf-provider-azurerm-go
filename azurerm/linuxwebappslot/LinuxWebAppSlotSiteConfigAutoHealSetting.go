package linuxwebappslot


type LinuxWebAppSlotSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/linux_web_app_slot#action LinuxWebAppSlot#action}
	Action *LinuxWebAppSlotSiteConfigAutoHealSettingAction `field:"optional" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/linux_web_app_slot#trigger LinuxWebAppSlot#trigger}
	Trigger *LinuxWebAppSlotSiteConfigAutoHealSettingTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

