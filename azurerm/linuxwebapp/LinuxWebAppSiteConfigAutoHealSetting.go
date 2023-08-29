// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_web_app#action LinuxWebApp#action}
	Action *LinuxWebAppSiteConfigAutoHealSettingAction `field:"optional" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_web_app#trigger LinuxWebApp#trigger}
	Trigger *LinuxWebAppSiteConfigAutoHealSettingTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

