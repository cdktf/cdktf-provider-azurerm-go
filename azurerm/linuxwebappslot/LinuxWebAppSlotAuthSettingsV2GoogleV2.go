package linuxwebappslot


type LinuxWebAppSlotAuthSettingsV2GoogleV2 struct {
	// The OpenID Connect Client ID for the Google web application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#client_id LinuxWebAppSlot#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The app setting name that contains the `client_secret` value used for Google Login.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#client_secret_setting_name LinuxWebAppSlot#client_secret_setting_name}
	ClientSecretSettingName *string `field:"required" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// Specifies a list of Allowed Audiences that will be requested as part of Google Sign-In authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#allowed_audiences LinuxWebAppSlot#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// Specifies a list of Login scopes that will be requested as part of Google Sign-In authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#login_scopes LinuxWebAppSlot#login_scopes}
	LoginScopes *[]*string `field:"optional" json:"loginScopes" yaml:"loginScopes"`
}
