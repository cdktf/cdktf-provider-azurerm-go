package linuxwebapp


type LinuxWebAppAuthSettingsV2AppleV2 struct {
	// The OpenID Connect Client ID for the Apple web application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_web_app#client_id LinuxWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The app setting name that contains the `client_secret` value used for Apple Login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_web_app#client_secret_setting_name LinuxWebApp#client_secret_setting_name}
	ClientSecretSettingName *string `field:"required" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
}

