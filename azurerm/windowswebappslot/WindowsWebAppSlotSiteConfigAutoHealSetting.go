// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/windows_web_app_slot#action WindowsWebAppSlot#action}
	Action *WindowsWebAppSlotSiteConfigAutoHealSettingAction `field:"required" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/windows_web_app_slot#trigger WindowsWebAppSlot#trigger}
	Trigger *WindowsWebAppSlotSiteConfigAutoHealSettingTrigger `field:"required" json:"trigger" yaml:"trigger"`
}

