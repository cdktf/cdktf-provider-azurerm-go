// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/windows_web_app_slot#count WindowsWebAppSlot#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/windows_web_app_slot#interval WindowsWebAppSlot#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
}

