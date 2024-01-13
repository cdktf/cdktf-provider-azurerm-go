// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktopworkspaceapplicationgroupassociation


type VirtualDesktopWorkspaceApplicationGroupAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/virtual_desktop_workspace_application_group_association#create VirtualDesktopWorkspaceApplicationGroupAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/virtual_desktop_workspace_application_group_association#delete VirtualDesktopWorkspaceApplicationGroupAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/virtual_desktop_workspace_application_group_association#read VirtualDesktopWorkspaceApplicationGroupAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

