// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devcenterprojectenvironmenttype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevCenterProjectEnvironmentTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#deployment_target_id DevCenterProjectEnvironmentType#deployment_target_id}.
	DeploymentTargetId *string `field:"required" json:"deploymentTargetId" yaml:"deploymentTargetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#dev_center_project_id DevCenterProjectEnvironmentType#dev_center_project_id}.
	DevCenterProjectId *string `field:"required" json:"devCenterProjectId" yaml:"devCenterProjectId"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#identity DevCenterProjectEnvironmentType#identity}
	Identity *DevCenterProjectEnvironmentTypeIdentity `field:"required" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#location DevCenterProjectEnvironmentType#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#name DevCenterProjectEnvironmentType#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#creator_role_assignment_roles DevCenterProjectEnvironmentType#creator_role_assignment_roles}.
	CreatorRoleAssignmentRoles *[]*string `field:"optional" json:"creatorRoleAssignmentRoles" yaml:"creatorRoleAssignmentRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#id DevCenterProjectEnvironmentType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#tags DevCenterProjectEnvironmentType#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#timeouts DevCenterProjectEnvironmentType#timeouts}
	Timeouts *DevCenterProjectEnvironmentTypeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user_role_assignment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/dev_center_project_environment_type#user_role_assignment DevCenterProjectEnvironmentType#user_role_assignment}
	UserRoleAssignment interface{} `field:"optional" json:"userRoleAssignment" yaml:"userRoleAssignment"`
}

