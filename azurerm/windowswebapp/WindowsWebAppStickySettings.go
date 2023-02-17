package windowswebapp


type WindowsWebAppStickySettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#app_setting_names WindowsWebApp#app_setting_names}.
	AppSettingNames *[]*string `field:"optional" json:"appSettingNames" yaml:"appSettingNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app#connection_string_names WindowsWebApp#connection_string_names}.
	ConnectionStringNames *[]*string `field:"optional" json:"connectionStringNames" yaml:"connectionStringNames"`
}

