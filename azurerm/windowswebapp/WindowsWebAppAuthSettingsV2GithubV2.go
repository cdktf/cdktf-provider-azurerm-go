// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppAuthSettingsV2GithubV2 struct {
	// The ID of the GitHub app used for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/windows_web_app#client_id WindowsWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The app setting name that contains the `client_secret` value used for GitHub Login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/windows_web_app#client_secret_setting_name WindowsWebApp#client_secret_setting_name}
	ClientSecretSettingName *string `field:"required" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// Specifies a list of OAuth 2.0 scopes that will be requested as part of GitHub Login authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/windows_web_app#login_scopes WindowsWebApp#login_scopes}
	LoginScopes *[]*string `field:"optional" json:"loginScopes" yaml:"loginScopes"`
}

