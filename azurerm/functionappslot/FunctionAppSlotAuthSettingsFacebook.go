package functionappslot


type FunctionAppSlotAuthSettingsFacebook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/function_app_slot#app_id FunctionAppSlot#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/function_app_slot#app_secret FunctionAppSlot#app_secret}.
	AppSecret *string `field:"required" json:"appSecret" yaml:"appSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/function_app_slot#oauth_scopes FunctionAppSlot#oauth_scopes}.
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

