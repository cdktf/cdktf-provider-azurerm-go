package linuxwebapp


type LinuxWebAppAuthSettingsV2TwitterV2 struct {
	// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#consumer_key LinuxWebApp#consumer_key}
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// The app setting name that contains the OAuth 1.0a consumer secret of the Twitter application used for sign-in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#consumer_secret_setting_name LinuxWebApp#consumer_secret_setting_name}
	ConsumerSecretSettingName *string `field:"required" json:"consumerSecretSettingName" yaml:"consumerSecretSettingName"`
}
