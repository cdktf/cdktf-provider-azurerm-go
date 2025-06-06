// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionapp


type FunctionAppAuthSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#enabled FunctionApp#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#active_directory FunctionApp#active_directory}
	ActiveDirectory *FunctionAppAuthSettingsActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#additional_login_params FunctionApp#additional_login_params}.
	AdditionalLoginParams *map[string]*string `field:"optional" json:"additionalLoginParams" yaml:"additionalLoginParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#allowed_external_redirect_urls FunctionApp#allowed_external_redirect_urls}.
	AllowedExternalRedirectUrls *[]*string `field:"optional" json:"allowedExternalRedirectUrls" yaml:"allowedExternalRedirectUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#default_provider FunctionApp#default_provider}.
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// facebook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#facebook FunctionApp#facebook}
	Facebook *FunctionAppAuthSettingsFacebook `field:"optional" json:"facebook" yaml:"facebook"`
	// google block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#google FunctionApp#google}
	Google *FunctionAppAuthSettingsGoogle `field:"optional" json:"google" yaml:"google"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#issuer FunctionApp#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// microsoft block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#microsoft FunctionApp#microsoft}
	Microsoft *FunctionAppAuthSettingsMicrosoft `field:"optional" json:"microsoft" yaml:"microsoft"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#runtime_version FunctionApp#runtime_version}.
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#token_refresh_extension_hours FunctionApp#token_refresh_extension_hours}.
	TokenRefreshExtensionHours *float64 `field:"optional" json:"tokenRefreshExtensionHours" yaml:"tokenRefreshExtensionHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#token_store_enabled FunctionApp#token_store_enabled}.
	TokenStoreEnabled interface{} `field:"optional" json:"tokenStoreEnabled" yaml:"tokenStoreEnabled"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#twitter FunctionApp#twitter}
	Twitter *FunctionAppAuthSettingsTwitter `field:"optional" json:"twitter" yaml:"twitter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/function_app#unauthenticated_client_action FunctionApp#unauthenticated_client_action}.
	UnauthenticatedClientAction *string `field:"optional" json:"unauthenticatedClientAction" yaml:"unauthenticatedClientAction"`
}

