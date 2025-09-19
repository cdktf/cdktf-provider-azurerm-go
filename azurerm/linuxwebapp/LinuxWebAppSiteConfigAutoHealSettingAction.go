// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSettingAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/linux_web_app#action_type LinuxWebApp#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/linux_web_app#minimum_process_execution_time LinuxWebApp#minimum_process_execution_time}.
	MinimumProcessExecutionTime *string `field:"optional" json:"minimumProcessExecutionTime" yaml:"minimumProcessExecutionTime"`
}

