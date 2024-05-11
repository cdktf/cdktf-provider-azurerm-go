// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot


type LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/linux_web_app_slot#count LinuxWebAppSlot#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/linux_web_app_slot#interval LinuxWebAppSlot#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/linux_web_app_slot#time_taken LinuxWebAppSlot#time_taken}.
	TimeTaken *string `field:"required" json:"timeTaken" yaml:"timeTaken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/linux_web_app_slot#path LinuxWebAppSlot#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

