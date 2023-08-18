package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSettingTriggerStatusCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/windows_web_app_slot#count WindowsWebAppSlot#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/windows_web_app_slot#interval WindowsWebAppSlot#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/windows_web_app_slot#status_code_range WindowsWebAppSlot#status_code_range}.
	StatusCodeRange *string `field:"required" json:"statusCodeRange" yaml:"statusCodeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/windows_web_app_slot#path WindowsWebAppSlot#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/windows_web_app_slot#sub_status WindowsWebAppSlot#sub_status}.
	SubStatus *float64 `field:"optional" json:"subStatus" yaml:"subStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/windows_web_app_slot#win32_status WindowsWebAppSlot#win32_status}.
	Win32Status *float64 `field:"optional" json:"win32Status" yaml:"win32Status"`
}

