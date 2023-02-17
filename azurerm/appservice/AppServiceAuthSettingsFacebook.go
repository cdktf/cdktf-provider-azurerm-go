package appservice


type AppServiceAuthSettingsFacebook struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#app_id AppService#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#app_secret AppService#app_secret}.
	AppSecret *string `field:"required" json:"appSecret" yaml:"appSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#oauth_scopes AppService#oauth_scopes}.
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

