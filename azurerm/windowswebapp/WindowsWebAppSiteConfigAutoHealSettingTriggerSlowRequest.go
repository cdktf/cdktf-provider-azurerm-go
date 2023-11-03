// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppSiteConfigAutoHealSettingTriggerSlowRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#count WindowsWebApp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#interval WindowsWebApp#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#time_taken WindowsWebApp#time_taken}.
	TimeTaken *string `field:"required" json:"timeTaken" yaml:"timeTaken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#path WindowsWebApp#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

