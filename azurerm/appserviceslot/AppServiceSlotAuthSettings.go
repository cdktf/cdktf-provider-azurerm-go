// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appserviceslot


type AppServiceSlotAuthSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#enabled AppServiceSlot#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#active_directory AppServiceSlot#active_directory}
	ActiveDirectory *AppServiceSlotAuthSettingsActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#additional_login_params AppServiceSlot#additional_login_params}.
	AdditionalLoginParams *map[string]*string `field:"optional" json:"additionalLoginParams" yaml:"additionalLoginParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#allowed_external_redirect_urls AppServiceSlot#allowed_external_redirect_urls}.
	AllowedExternalRedirectUrls *[]*string `field:"optional" json:"allowedExternalRedirectUrls" yaml:"allowedExternalRedirectUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#default_provider AppServiceSlot#default_provider}.
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// facebook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#facebook AppServiceSlot#facebook}
	Facebook *AppServiceSlotAuthSettingsFacebook `field:"optional" json:"facebook" yaml:"facebook"`
	// google block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#google AppServiceSlot#google}
	Google *AppServiceSlotAuthSettingsGoogle `field:"optional" json:"google" yaml:"google"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#issuer AppServiceSlot#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// microsoft block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#microsoft AppServiceSlot#microsoft}
	Microsoft *AppServiceSlotAuthSettingsMicrosoft `field:"optional" json:"microsoft" yaml:"microsoft"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#runtime_version AppServiceSlot#runtime_version}.
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#token_refresh_extension_hours AppServiceSlot#token_refresh_extension_hours}.
	TokenRefreshExtensionHours *float64 `field:"optional" json:"tokenRefreshExtensionHours" yaml:"tokenRefreshExtensionHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#token_store_enabled AppServiceSlot#token_store_enabled}.
	TokenStoreEnabled interface{} `field:"optional" json:"tokenStoreEnabled" yaml:"tokenStoreEnabled"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#twitter AppServiceSlot#twitter}
	Twitter *AppServiceSlotAuthSettingsTwitter `field:"optional" json:"twitter" yaml:"twitter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service_slot#unauthenticated_client_action AppServiceSlot#unauthenticated_client_action}.
	UnauthenticatedClientAction *string `field:"optional" json:"unauthenticatedClientAction" yaml:"unauthenticatedClientAction"`
}

