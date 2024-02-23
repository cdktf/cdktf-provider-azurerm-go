// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionapp


type WindowsFunctionAppAuthSettingsMicrosoft struct {
	// The OAuth 2.0 client ID that was created for the app used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/windows_function_app#client_id WindowsFunctionApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret that was created for the app used for authentication. Cannot be specified with `client_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/windows_function_app#client_secret WindowsFunctionApp#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The app setting name containing the OAuth 2.0 client secret that was created for the app used for authentication. Cannot be specified with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/windows_function_app#client_secret_setting_name WindowsFunctionApp#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// The list of OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication. If not specified, `wl.basic` is used as the default scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/windows_function_app#oauth_scopes WindowsFunctionApp#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

