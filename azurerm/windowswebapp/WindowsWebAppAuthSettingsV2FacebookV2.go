// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppAuthSettingsV2FacebookV2 struct {
	// The App ID of the Facebook app used for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#app_id WindowsWebApp#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The app setting name that contains the `app_secret` value used for Facebook Login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#app_secret_setting_name WindowsWebApp#app_secret_setting_name}
	AppSecretSettingName *string `field:"required" json:"appSecretSettingName" yaml:"appSecretSettingName"`
	// The version of the Facebook API to be used while logging in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#graph_api_version WindowsWebApp#graph_api_version}
	GraphApiVersion *string `field:"optional" json:"graphApiVersion" yaml:"graphApiVersion"`
	// Specifies a list of scopes to be requested as part of Facebook Login authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#login_scopes WindowsWebApp#login_scopes}
	LoginScopes *[]*string `field:"optional" json:"loginScopes" yaml:"loginScopes"`
}

