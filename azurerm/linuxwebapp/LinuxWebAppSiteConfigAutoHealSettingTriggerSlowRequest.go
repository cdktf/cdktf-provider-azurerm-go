package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequest struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#count LinuxWebApp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#interval LinuxWebApp#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#time_taken LinuxWebApp#time_taken}.
	TimeTaken *string `field:"required" json:"timeTaken" yaml:"timeTaken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#path LinuxWebApp#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}
