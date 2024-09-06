// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionapp


type WindowsFunctionAppAuthSettings struct {
	// Should the Authentication / Authorization feature be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#enabled WindowsFunctionApp#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#active_directory WindowsFunctionApp#active_directory}
	ActiveDirectory *WindowsFunctionAppAuthSettingsActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Specifies a map of Login Parameters to send to the OpenID Connect authorization endpoint when a user logs in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#additional_login_parameters WindowsFunctionApp#additional_login_parameters}
	AdditionalLoginParameters *map[string]*string `field:"optional" json:"additionalLoginParameters" yaml:"additionalLoginParameters"`
	// Specifies a list of External URLs that can be redirected to as part of logging in or logging out of the Windows Web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#allowed_external_redirect_urls WindowsFunctionApp#allowed_external_redirect_urls}
	AllowedExternalRedirectUrls *[]*string `field:"optional" json:"allowedExternalRedirectUrls" yaml:"allowedExternalRedirectUrls"`
	// The default authentication provider to use when multiple providers are configured.
	//
	// Possible values include: `AzureActiveDirectory`, `Facebook`, `Google`, `MicrosoftAccount`, `Twitter`, `Github`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#default_provider WindowsFunctionApp#default_provider}
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// facebook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#facebook WindowsFunctionApp#facebook}
	Facebook *WindowsFunctionAppAuthSettingsFacebook `field:"optional" json:"facebook" yaml:"facebook"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#github WindowsFunctionApp#github}
	Github *WindowsFunctionAppAuthSettingsGithub `field:"optional" json:"github" yaml:"github"`
	// google block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#google WindowsFunctionApp#google}
	Google *WindowsFunctionAppAuthSettingsGoogle `field:"optional" json:"google" yaml:"google"`
	// The OpenID Connect Issuer URI that represents the entity which issues access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#issuer WindowsFunctionApp#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// microsoft block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#microsoft WindowsFunctionApp#microsoft}
	Microsoft *WindowsFunctionAppAuthSettingsMicrosoft `field:"optional" json:"microsoft" yaml:"microsoft"`
	// The RuntimeVersion of the Authentication / Authorization feature in use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#runtime_version WindowsFunctionApp#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// The number of hours after session token expiration that a session token can be used to call the token refresh API.
	//
	// Defaults to `72` hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#token_refresh_extension_hours WindowsFunctionApp#token_refresh_extension_hours}
	TokenRefreshExtensionHours *float64 `field:"optional" json:"tokenRefreshExtensionHours" yaml:"tokenRefreshExtensionHours"`
	// Should the Windows Web App durably store platform-specific security tokens that are obtained during login flows? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#token_store_enabled WindowsFunctionApp#token_store_enabled}
	TokenStoreEnabled interface{} `field:"optional" json:"tokenStoreEnabled" yaml:"tokenStoreEnabled"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#twitter WindowsFunctionApp#twitter}
	Twitter *WindowsFunctionAppAuthSettingsTwitter `field:"optional" json:"twitter" yaml:"twitter"`
	// The action to take when an unauthenticated client attempts to access the app. Possible values include: `RedirectToLoginPage`, `AllowAnonymous`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/windows_function_app#unauthenticated_client_action WindowsFunctionApp#unauthenticated_client_action}
	UnauthenticatedClientAction *string `field:"optional" json:"unauthenticatedClientAction" yaml:"unauthenticatedClientAction"`
}

