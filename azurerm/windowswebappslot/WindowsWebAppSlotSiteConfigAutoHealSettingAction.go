// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSettingAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/windows_web_app_slot#action_type WindowsWebAppSlot#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// custom_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/windows_web_app_slot#custom_action WindowsWebAppSlot#custom_action}
	CustomAction *WindowsWebAppSlotSiteConfigAutoHealSettingActionCustomAction `field:"optional" json:"customAction" yaml:"customAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/windows_web_app_slot#minimum_process_execution_time WindowsWebAppSlot#minimum_process_execution_time}.
	MinimumProcessExecutionTime *string `field:"optional" json:"minimumProcessExecutionTime" yaml:"minimumProcessExecutionTime"`
}

