package windowsfunctionapp


type WindowsFunctionAppStickySettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_function_app#app_setting_names WindowsFunctionApp#app_setting_names}.
	AppSettingNames *[]*string `field:"optional" json:"appSettingNames" yaml:"appSettingNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_function_app#connection_string_names WindowsFunctionApp#connection_string_names}.
	ConnectionStringNames *[]*string `field:"optional" json:"connectionStringNames" yaml:"connectionStringNames"`
}
