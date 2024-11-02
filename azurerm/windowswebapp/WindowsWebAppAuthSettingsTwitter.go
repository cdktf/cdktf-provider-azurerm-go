// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppAuthSettingsTwitter struct {
	// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/windows_web_app#consumer_key WindowsWebApp#consumer_key}
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// The OAuth 1.0a consumer secret of the Twitter application used for sign-in. Cannot be specified with `consumer_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/windows_web_app#consumer_secret WindowsWebApp#consumer_secret}
	ConsumerSecret *string `field:"optional" json:"consumerSecret" yaml:"consumerSecret"`
	// The app setting name that contains the OAuth 1.0a consumer secret of the Twitter application used for sign-in. Cannot be specified with `consumer_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/windows_web_app#consumer_secret_setting_name WindowsWebApp#consumer_secret_setting_name}
	ConsumerSecretSettingName *string `field:"optional" json:"consumerSecretSettingName" yaml:"consumerSecretSettingName"`
}

