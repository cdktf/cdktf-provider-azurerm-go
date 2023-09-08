// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/linux_web_app#count LinuxWebApp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/linux_web_app#interval LinuxWebApp#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/linux_web_app#time_taken LinuxWebApp#time_taken}.
	TimeTaken *string `field:"required" json:"timeTaken" yaml:"timeTaken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/linux_web_app#path LinuxWebApp#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

