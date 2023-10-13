// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppSiteConfigVirtualApplicationVirtualDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/windows_web_app#physical_path WindowsWebApp#physical_path}.
	PhysicalPath *string `field:"optional" json:"physicalPath" yaml:"physicalPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/windows_web_app#virtual_path WindowsWebApp#virtual_path}.
	VirtualPath *string `field:"optional" json:"virtualPath" yaml:"virtualPath"`
}

