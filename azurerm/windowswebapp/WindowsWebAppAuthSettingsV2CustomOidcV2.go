package windowswebapp


type WindowsWebAppAuthSettingsV2CustomOidcV2 struct {
	// The ID of the Client to use to authenticate with this Custom OIDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/windows_web_app#client_id WindowsWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The name of the Custom OIDC Authentication Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/windows_web_app#name WindowsWebApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The endpoint that contains all the configuration endpoints for this Custom OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/windows_web_app#openid_configuration_endpoint WindowsWebApp#openid_configuration_endpoint}
	OpenidConfigurationEndpoint *string `field:"required" json:"openidConfigurationEndpoint" yaml:"openidConfigurationEndpoint"`
	// The name of the claim that contains the users name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/windows_web_app#name_claim_type WindowsWebApp#name_claim_type}
	NameClaimType *string `field:"optional" json:"nameClaimType" yaml:"nameClaimType"`
	// The list of the scopes that should be requested while authenticating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/windows_web_app#scopes WindowsWebApp#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

