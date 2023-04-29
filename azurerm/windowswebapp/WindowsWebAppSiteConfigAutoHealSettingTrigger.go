package windowswebapp


type WindowsWebAppSiteConfigAutoHealSettingTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/windows_web_app#private_memory_kb WindowsWebApp#private_memory_kb}.
	PrivateMemoryKb *float64 `field:"optional" json:"privateMemoryKb" yaml:"privateMemoryKb"`
	// requests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/windows_web_app#requests WindowsWebApp#requests}
	Requests *WindowsWebAppSiteConfigAutoHealSettingTriggerRequests `field:"optional" json:"requests" yaml:"requests"`
	// slow_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/windows_web_app#slow_request WindowsWebApp#slow_request}
	SlowRequest *WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest `field:"optional" json:"slowRequest" yaml:"slowRequest"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/windows_web_app#status_code WindowsWebApp#status_code}
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

