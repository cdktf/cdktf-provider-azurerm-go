// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp


type LinuxFunctionAppStickySettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/linux_function_app#app_setting_names LinuxFunctionApp#app_setting_names}.
	AppSettingNames *[]*string `field:"optional" json:"appSettingNames" yaml:"appSettingNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/linux_function_app#connection_string_names LinuxFunctionApp#connection_string_names}.
	ConnectionStringNames *[]*string `field:"optional" json:"connectionStringNames" yaml:"connectionStringNames"`
}

