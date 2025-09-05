// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionStickySettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/function_app_flex_consumption#app_setting_names FunctionAppFlexConsumption#app_setting_names}.
	AppSettingNames *[]*string `field:"optional" json:"appSettingNames" yaml:"appSettingNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/function_app_flex_consumption#connection_string_names FunctionAppFlexConsumption#connection_string_names}.
	ConnectionStringNames *[]*string `field:"optional" json:"connectionStringNames" yaml:"connectionStringNames"`
}

