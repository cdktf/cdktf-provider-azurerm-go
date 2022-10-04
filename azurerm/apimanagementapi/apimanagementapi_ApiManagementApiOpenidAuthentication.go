package apimanagementapi


type ApiManagementApiOpenidAuthentication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api#openid_provider_name ApiManagementApi#openid_provider_name}.
	OpenidProviderName *string `field:"required" json:"openidProviderName" yaml:"openidProviderName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_api#bearer_token_sending_methods ApiManagementApi#bearer_token_sending_methods}.
	BearerTokenSendingMethods *[]*string `field:"optional" json:"bearerTokenSendingMethods" yaml:"bearerTokenSendingMethods"`
}

