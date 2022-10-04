package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#action LinuxWebApp#action}
	Action *LinuxWebAppSiteConfigAutoHealSettingAction `field:"optional" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#trigger LinuxWebApp#trigger}
	Trigger *LinuxWebAppSiteConfigAutoHealSettingTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

