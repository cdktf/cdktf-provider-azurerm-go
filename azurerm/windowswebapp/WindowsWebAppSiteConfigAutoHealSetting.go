// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/windows_web_app#action WindowsWebApp#action}
	Action *WindowsWebAppSiteConfigAutoHealSettingAction `field:"required" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/windows_web_app#trigger WindowsWebApp#trigger}
	Trigger *WindowsWebAppSiteConfigAutoHealSettingTrigger `field:"required" json:"trigger" yaml:"trigger"`
}

