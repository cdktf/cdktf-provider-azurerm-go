package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSettingTriggerStatusCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/linux_web_app#count LinuxWebApp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/linux_web_app#interval LinuxWebApp#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/linux_web_app#status_code_range LinuxWebApp#status_code_range}.
	StatusCodeRange *string `field:"required" json:"statusCodeRange" yaml:"statusCodeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/linux_web_app#path LinuxWebApp#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/linux_web_app#sub_status LinuxWebApp#sub_status}.
	SubStatus *float64 `field:"optional" json:"subStatus" yaml:"subStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/linux_web_app#win32_status LinuxWebApp#win32_status}.
	Win32Status *float64 `field:"optional" json:"win32Status" yaml:"win32Status"`
}

