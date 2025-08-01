// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot


type LinuxWebAppSlotSiteConfigAutoHealSettingAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_web_app_slot#action_type LinuxWebAppSlot#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/linux_web_app_slot#minimum_process_execution_time LinuxWebAppSlot#minimum_process_execution_time}.
	MinimumProcessExecutionTime *string `field:"optional" json:"minimumProcessExecutionTime" yaml:"minimumProcessExecutionTime"`
}

