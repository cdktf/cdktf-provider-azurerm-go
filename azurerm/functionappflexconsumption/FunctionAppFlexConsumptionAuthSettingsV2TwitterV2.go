// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionAuthSettingsV2TwitterV2 struct {
	// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/function_app_flex_consumption#consumer_key FunctionAppFlexConsumption#consumer_key}
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// The app setting name that contains the OAuth 1.0a consumer secret of the Twitter application used for sign-in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/function_app_flex_consumption#consumer_secret_setting_name FunctionAppFlexConsumption#consumer_secret_setting_name}
	ConsumerSecretSettingName *string `field:"required" json:"consumerSecretSettingName" yaml:"consumerSecretSettingName"`
}

