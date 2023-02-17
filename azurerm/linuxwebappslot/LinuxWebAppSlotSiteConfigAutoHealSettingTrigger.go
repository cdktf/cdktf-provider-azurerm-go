package linuxwebappslot


type LinuxWebAppSlotSiteConfigAutoHealSettingTrigger struct {
	// requests block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#requests LinuxWebAppSlot#requests}
	Requests *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequests `field:"optional" json:"requests" yaml:"requests"`
	// slow_request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#slow_request LinuxWebAppSlot#slow_request}
	SlowRequest *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest `field:"optional" json:"slowRequest" yaml:"slowRequest"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#status_code LinuxWebAppSlot#status_code}
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

