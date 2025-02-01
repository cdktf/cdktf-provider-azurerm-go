// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimactiveroleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PimActiveRoleAssignmentConfig struct {
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
	// Object ID of the principal for this role assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/pim_active_role_assignment#principal_id PimActiveRoleAssignment#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Role definition ID for this role assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/pim_active_role_assignment#role_definition_id PimActiveRoleAssignment#role_definition_id}
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// Scope for this role assignment, should be a valid resource ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/pim_active_role_assignment#scope PimActiveRoleAssignment#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/pim_active_role_assignment#id PimActiveRoleAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The justification for this role assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/pim_active_role_assignment#justification PimActiveRoleAssignment#justification}
	Justification *string `field:"optional" json:"justification" yaml:"justification"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/pim_active_role_assignment#schedule PimActiveRoleAssignment#schedule}
	Schedule *PimActiveRoleAssignmentSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// ticket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/pim_active_role_assignment#ticket PimActiveRoleAssignment#ticket}
	Ticket *PimActiveRoleAssignmentTicket `field:"optional" json:"ticket" yaml:"ticket"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/pim_active_role_assignment#timeouts PimActiveRoleAssignment#timeouts}
	Timeouts *PimActiveRoleAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

