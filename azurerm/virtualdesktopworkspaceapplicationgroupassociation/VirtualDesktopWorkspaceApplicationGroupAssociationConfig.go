// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktopworkspaceapplicationgroupassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualDesktopWorkspaceApplicationGroupAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/virtual_desktop_workspace_application_group_association#application_group_id VirtualDesktopWorkspaceApplicationGroupAssociation#application_group_id}.
	ApplicationGroupId *string `field:"required" json:"applicationGroupId" yaml:"applicationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/virtual_desktop_workspace_application_group_association#workspace_id VirtualDesktopWorkspaceApplicationGroupAssociation#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/virtual_desktop_workspace_application_group_association#id VirtualDesktopWorkspaceApplicationGroupAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/virtual_desktop_workspace_application_group_association#timeouts VirtualDesktopWorkspaceApplicationGroupAssociation#timeouts}
	Timeouts *VirtualDesktopWorkspaceApplicationGroupAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

