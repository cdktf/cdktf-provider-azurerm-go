package windowsfunctionapp


type WindowsFunctionAppAuthSettingsV2CustomOidcV2 struct {
	// The ID of the Client to use to authenticate with this Custom OIDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#client_id WindowsFunctionApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The name of the Custom OIDC Authentication Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#name WindowsFunctionApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The endpoint that contains all the configuration endpoints for this Custom OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#openid_configuration_endpoint WindowsFunctionApp#openid_configuration_endpoint}
	OpenidConfigurationEndpoint *string `field:"required" json:"openidConfigurationEndpoint" yaml:"openidConfigurationEndpoint"`
	// The name of the claim that contains the users name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#name_claim_type WindowsFunctionApp#name_claim_type}
	NameClaimType *string `field:"optional" json:"nameClaimType" yaml:"nameClaimType"`
	// The list of the scopes that should be requested while authenticating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#scopes WindowsFunctionApp#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

