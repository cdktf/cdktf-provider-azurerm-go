package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSettingActionCustomAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#executable WindowsWebAppSlot#executable}.
	Executable *string `field:"required" json:"executable" yaml:"executable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#parameters WindowsWebAppSlot#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}
