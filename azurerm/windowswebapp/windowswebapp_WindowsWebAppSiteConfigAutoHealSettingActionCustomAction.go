package windowswebapp


type WindowsWebAppSiteConfigAutoHealSettingActionCustomAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#executable WindowsWebApp#executable}.
	Executable *string `field:"required" json:"executable" yaml:"executable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#parameters WindowsWebApp#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

