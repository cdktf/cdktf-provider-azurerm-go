// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktopapplicationgroup


type VirtualDesktopApplicationGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_desktop_application_group#create VirtualDesktopApplicationGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_desktop_application_group#delete VirtualDesktopApplicationGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_desktop_application_group#read VirtualDesktopApplicationGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_desktop_application_group#update VirtualDesktopApplicationGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

