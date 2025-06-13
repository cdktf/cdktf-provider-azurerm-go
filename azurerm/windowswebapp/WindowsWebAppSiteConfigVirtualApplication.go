// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppSiteConfigVirtualApplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/windows_web_app#physical_path WindowsWebApp#physical_path}.
	PhysicalPath *string `field:"required" json:"physicalPath" yaml:"physicalPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/windows_web_app#preload WindowsWebApp#preload}.
	Preload interface{} `field:"required" json:"preload" yaml:"preload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/windows_web_app#virtual_path WindowsWebApp#virtual_path}.
	VirtualPath *string `field:"required" json:"virtualPath" yaml:"virtualPath"`
	// virtual_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/windows_web_app#virtual_directory WindowsWebApp#virtual_directory}
	VirtualDirectory interface{} `field:"optional" json:"virtualDirectory" yaml:"virtualDirectory"`
}

