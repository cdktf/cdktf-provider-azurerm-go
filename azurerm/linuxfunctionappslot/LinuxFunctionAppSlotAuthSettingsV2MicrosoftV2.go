// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionappslot


type LinuxFunctionAppSlotAuthSettingsV2MicrosoftV2 struct {
	// The OAuth 2.0 client ID that was created for the app used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_function_app_slot#client_id LinuxFunctionAppSlot#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The app setting name containing the OAuth 2.0 client secret that was created for the app used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_function_app_slot#client_secret_setting_name LinuxFunctionAppSlot#client_secret_setting_name}
	ClientSecretSettingName *string `field:"required" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// Specifies a list of Allowed Audiences that will be requested as part of Microsoft Sign-In authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_function_app_slot#allowed_audiences LinuxFunctionAppSlot#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// The list of Login scopes that will be requested as part of Microsoft Account authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_function_app_slot#login_scopes LinuxFunctionAppSlot#login_scopes}
	LoginScopes *[]*string `field:"optional" json:"loginScopes" yaml:"loginScopes"`
}

