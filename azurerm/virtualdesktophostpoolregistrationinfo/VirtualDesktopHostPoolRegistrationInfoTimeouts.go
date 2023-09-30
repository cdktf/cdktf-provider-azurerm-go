// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktophostpoolregistrationinfo


type VirtualDesktopHostPoolRegistrationInfoTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/virtual_desktop_host_pool_registration_info#create VirtualDesktopHostPoolRegistrationInfo#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/virtual_desktop_host_pool_registration_info#delete VirtualDesktopHostPoolRegistrationInfo#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/virtual_desktop_host_pool_registration_info#read VirtualDesktopHostPoolRegistrationInfo#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/virtual_desktop_host_pool_registration_info#update VirtualDesktopHostPoolRegistrationInfo#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

