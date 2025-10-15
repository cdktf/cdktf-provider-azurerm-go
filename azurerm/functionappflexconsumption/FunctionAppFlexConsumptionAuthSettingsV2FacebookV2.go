// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionAuthSettingsV2FacebookV2 struct {
	// The App ID of the Facebook app used for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/function_app_flex_consumption#app_id FunctionAppFlexConsumption#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The app setting name that contains the `app_secret` value used for Facebook Login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/function_app_flex_consumption#app_secret_setting_name FunctionAppFlexConsumption#app_secret_setting_name}
	AppSecretSettingName *string `field:"required" json:"appSecretSettingName" yaml:"appSecretSettingName"`
	// The version of the Facebook API to be used while logging in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/function_app_flex_consumption#graph_api_version FunctionAppFlexConsumption#graph_api_version}
	GraphApiVersion *string `field:"optional" json:"graphApiVersion" yaml:"graphApiVersion"`
	// Specifies a list of scopes to be requested as part of Facebook Login authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/function_app_flex_consumption#login_scopes FunctionAppFlexConsumption#login_scopes}
	LoginScopes *[]*string `field:"optional" json:"loginScopes" yaml:"loginScopes"`
}

