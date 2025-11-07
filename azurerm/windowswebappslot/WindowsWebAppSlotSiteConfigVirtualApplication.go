// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot


type WindowsWebAppSlotSiteConfigVirtualApplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/windows_web_app_slot#physical_path WindowsWebAppSlot#physical_path}.
	PhysicalPath *string `field:"required" json:"physicalPath" yaml:"physicalPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/windows_web_app_slot#preload WindowsWebAppSlot#preload}.
	Preload interface{} `field:"required" json:"preload" yaml:"preload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/windows_web_app_slot#virtual_path WindowsWebAppSlot#virtual_path}.
	VirtualPath *string `field:"required" json:"virtualPath" yaml:"virtualPath"`
	// virtual_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/windows_web_app_slot#virtual_directory WindowsWebAppSlot#virtual_directory}
	VirtualDirectory interface{} `field:"optional" json:"virtualDirectory" yaml:"virtualDirectory"`
}

