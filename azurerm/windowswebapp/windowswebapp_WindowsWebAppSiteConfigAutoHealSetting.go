package windowswebapp


type WindowsWebAppSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#action WindowsWebApp#action}
	Action *WindowsWebAppSiteConfigAutoHealSettingAction `field:"required" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#trigger WindowsWebApp#trigger}
	Trigger *WindowsWebAppSiteConfigAutoHealSettingTrigger `field:"required" json:"trigger" yaml:"trigger"`
}

