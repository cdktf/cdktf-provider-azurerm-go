// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleManagementPolicyConfig struct {
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
	// ID of the Azure Role to which this policy is assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/role_management_policy#role_definition_id RoleManagementPolicy#role_definition_id}
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// The scope of the role to which this policy will apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/role_management_policy#scope RoleManagementPolicy#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// activation_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/role_management_policy#activation_rules RoleManagementPolicy#activation_rules}
	ActivationRules *RoleManagementPolicyActivationRules `field:"optional" json:"activationRules" yaml:"activationRules"`
	// active_assignment_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/role_management_policy#active_assignment_rules RoleManagementPolicy#active_assignment_rules}
	ActiveAssignmentRules *RoleManagementPolicyActiveAssignmentRules `field:"optional" json:"activeAssignmentRules" yaml:"activeAssignmentRules"`
	// eligible_assignment_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/role_management_policy#eligible_assignment_rules RoleManagementPolicy#eligible_assignment_rules}
	EligibleAssignmentRules *RoleManagementPolicyEligibleAssignmentRules `field:"optional" json:"eligibleAssignmentRules" yaml:"eligibleAssignmentRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/role_management_policy#id RoleManagementPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// notification_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/role_management_policy#notification_rules RoleManagementPolicy#notification_rules}
	NotificationRules *RoleManagementPolicyNotificationRules `field:"optional" json:"notificationRules" yaml:"notificationRules"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/role_management_policy#timeouts RoleManagementPolicy#timeouts}
	Timeouts *RoleManagementPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

