// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot


type WindowsWebAppSlotSiteConfigHandlerMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/windows_web_app_slot#extension WindowsWebAppSlot#extension}.
	Extension *string `field:"required" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/windows_web_app_slot#script_processor_path WindowsWebAppSlot#script_processor_path}.
	ScriptProcessorPath *string `field:"required" json:"scriptProcessorPath" yaml:"scriptProcessorPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/windows_web_app_slot#arguments WindowsWebAppSlot#arguments}.
	Arguments *string `field:"optional" json:"arguments" yaml:"arguments"`
}

