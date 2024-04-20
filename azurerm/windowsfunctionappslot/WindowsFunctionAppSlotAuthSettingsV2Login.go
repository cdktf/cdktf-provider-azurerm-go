// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot


type WindowsFunctionAppSlotAuthSettingsV2Login struct {
	// External URLs that can be redirected to as part of logging in or logging out of the app.
	//
	// This is an advanced setting typically only needed by Windows Store application backends. **Note:** URLs within the current domain are always implicitly allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#allowed_external_redirect_urls WindowsFunctionAppSlot#allowed_external_redirect_urls}
	AllowedExternalRedirectUrls *[]*string `field:"optional" json:"allowedExternalRedirectUrls" yaml:"allowedExternalRedirectUrls"`
	// The method by which cookies expire. Possible values include: `FixedTime`, and `IdentityProviderDerived`. Defaults to `FixedTime`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#cookie_expiration_convention WindowsFunctionAppSlot#cookie_expiration_convention}
	CookieExpirationConvention *string `field:"optional" json:"cookieExpirationConvention" yaml:"cookieExpirationConvention"`
	// The time after the request is made when the session cookie should expire. Defaults to `08:00:00`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#cookie_expiration_time WindowsFunctionAppSlot#cookie_expiration_time}
	CookieExpirationTime *string `field:"optional" json:"cookieExpirationTime" yaml:"cookieExpirationTime"`
	// The endpoint to which logout requests should be made.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#logout_endpoint WindowsFunctionAppSlot#logout_endpoint}
	LogoutEndpoint *string `field:"optional" json:"logoutEndpoint" yaml:"logoutEndpoint"`
	// The time after the request is made when the nonce should expire. Defaults to `00:05:00`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#nonce_expiration_time WindowsFunctionAppSlot#nonce_expiration_time}
	NonceExpirationTime *string `field:"optional" json:"nonceExpirationTime" yaml:"nonceExpirationTime"`
	// Should the fragments from the request be preserved after the login request is made. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#preserve_url_fragments_for_logins WindowsFunctionAppSlot#preserve_url_fragments_for_logins}
	PreserveUrlFragmentsForLogins interface{} `field:"optional" json:"preserveUrlFragmentsForLogins" yaml:"preserveUrlFragmentsForLogins"`
	// The number of hours after session token expiration that a session token can be used to call the token refresh API.
	//
	// Defaults to `72` hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#token_refresh_extension_time WindowsFunctionAppSlot#token_refresh_extension_time}
	TokenRefreshExtensionTime *float64 `field:"optional" json:"tokenRefreshExtensionTime" yaml:"tokenRefreshExtensionTime"`
	// Should the Token Store configuration Enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#token_store_enabled WindowsFunctionAppSlot#token_store_enabled}
	TokenStoreEnabled interface{} `field:"optional" json:"tokenStoreEnabled" yaml:"tokenStoreEnabled"`
	// The directory path in the App Filesystem in which the tokens will be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#token_store_path WindowsFunctionAppSlot#token_store_path}
	TokenStorePath *string `field:"optional" json:"tokenStorePath" yaml:"tokenStorePath"`
	// The name of the app setting which contains the SAS URL of the blob storage containing the tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#token_store_sas_setting_name WindowsFunctionAppSlot#token_store_sas_setting_name}
	TokenStoreSasSettingName *string `field:"optional" json:"tokenStoreSasSettingName" yaml:"tokenStoreSasSettingName"`
	// Should the nonce be validated while completing the login flow. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/windows_function_app_slot#validate_nonce WindowsFunctionAppSlot#validate_nonce}
	ValidateNonce interface{} `field:"optional" json:"validateNonce" yaml:"validateNonce"`
}

