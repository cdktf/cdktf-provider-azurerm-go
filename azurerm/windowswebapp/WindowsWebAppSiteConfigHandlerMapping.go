// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppSiteConfigHandlerMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/windows_web_app#extension WindowsWebApp#extension}.
	Extension *string `field:"required" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/windows_web_app#script_processor_path WindowsWebApp#script_processor_path}.
	ScriptProcessorPath *string `field:"required" json:"scriptProcessorPath" yaml:"scriptProcessorPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/windows_web_app#arguments WindowsWebApp#arguments}.
	Arguments *string `field:"optional" json:"arguments" yaml:"arguments"`
}

