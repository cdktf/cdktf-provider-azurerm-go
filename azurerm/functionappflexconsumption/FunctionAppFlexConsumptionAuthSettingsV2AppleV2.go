// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionAuthSettingsV2AppleV2 struct {
	// The OpenID Connect Client ID for the Apple web application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#client_id FunctionAppFlexConsumption#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The app setting name that contains the `client_secret` value used for Apple Login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#client_secret_setting_name FunctionAppFlexConsumption#client_secret_setting_name}
	ClientSecretSettingName *string `field:"required" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
}

