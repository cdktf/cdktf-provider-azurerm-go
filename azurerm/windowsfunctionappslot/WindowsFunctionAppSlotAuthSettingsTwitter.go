// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot


type WindowsFunctionAppSlotAuthSettingsTwitter struct {
	// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/windows_function_app_slot#consumer_key WindowsFunctionAppSlot#consumer_key}
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// The OAuth 1.0a consumer secret of the Twitter application used for sign-in. Cannot be specified with `consumer_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/windows_function_app_slot#consumer_secret WindowsFunctionAppSlot#consumer_secret}
	ConsumerSecret *string `field:"optional" json:"consumerSecret" yaml:"consumerSecret"`
	// The app setting name that contains the OAuth 1.0a consumer secret of the Twitter application used for sign-in. Cannot be specified with `consumer_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/windows_function_app_slot#consumer_secret_setting_name WindowsFunctionAppSlot#consumer_secret_setting_name}
	ConsumerSecretSettingName *string `field:"optional" json:"consumerSecretSettingName" yaml:"consumerSecretSettingName"`
}

