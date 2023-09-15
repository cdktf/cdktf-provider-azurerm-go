// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppSiteConfigAutoHealSettingAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/windows_web_app#action_type WindowsWebApp#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// custom_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/windows_web_app#custom_action WindowsWebApp#custom_action}
	CustomAction *WindowsWebAppSiteConfigAutoHealSettingActionCustomAction `field:"optional" json:"customAction" yaml:"customAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/windows_web_app#minimum_process_execution_time WindowsWebApp#minimum_process_execution_time}.
	MinimumProcessExecutionTime *string `field:"optional" json:"minimumProcessExecutionTime" yaml:"minimumProcessExecutionTime"`
}

