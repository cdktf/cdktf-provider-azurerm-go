// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappslot


type FunctionAppSlotAuthSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#enabled FunctionAppSlot#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#active_directory FunctionAppSlot#active_directory}
	ActiveDirectory *FunctionAppSlotAuthSettingsActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#additional_login_params FunctionAppSlot#additional_login_params}.
	AdditionalLoginParams *map[string]*string `field:"optional" json:"additionalLoginParams" yaml:"additionalLoginParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#allowed_external_redirect_urls FunctionAppSlot#allowed_external_redirect_urls}.
	AllowedExternalRedirectUrls *[]*string `field:"optional" json:"allowedExternalRedirectUrls" yaml:"allowedExternalRedirectUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#default_provider FunctionAppSlot#default_provider}.
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// facebook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#facebook FunctionAppSlot#facebook}
	Facebook *FunctionAppSlotAuthSettingsFacebook `field:"optional" json:"facebook" yaml:"facebook"`
	// google block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#google FunctionAppSlot#google}
	Google *FunctionAppSlotAuthSettingsGoogle `field:"optional" json:"google" yaml:"google"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#issuer FunctionAppSlot#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// microsoft block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#microsoft FunctionAppSlot#microsoft}
	Microsoft *FunctionAppSlotAuthSettingsMicrosoft `field:"optional" json:"microsoft" yaml:"microsoft"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#runtime_version FunctionAppSlot#runtime_version}.
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#token_refresh_extension_hours FunctionAppSlot#token_refresh_extension_hours}.
	TokenRefreshExtensionHours *float64 `field:"optional" json:"tokenRefreshExtensionHours" yaml:"tokenRefreshExtensionHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#token_store_enabled FunctionAppSlot#token_store_enabled}.
	TokenStoreEnabled interface{} `field:"optional" json:"tokenStoreEnabled" yaml:"tokenStoreEnabled"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#twitter FunctionAppSlot#twitter}
	Twitter *FunctionAppSlotAuthSettingsTwitter `field:"optional" json:"twitter" yaml:"twitter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/function_app_slot#unauthenticated_client_action FunctionAppSlot#unauthenticated_client_action}.
	UnauthenticatedClientAction *string `field:"optional" json:"unauthenticatedClientAction" yaml:"unauthenticatedClientAction"`
}

