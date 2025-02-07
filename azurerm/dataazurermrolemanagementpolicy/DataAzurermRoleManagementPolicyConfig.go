// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermrolemanagementpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermRoleManagementPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/role_management_policy#role_definition_id DataAzurermRoleManagementPolicy#role_definition_id}
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// The scope of the role to which this policy will apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/role_management_policy#scope DataAzurermRoleManagementPolicy#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/role_management_policy#id DataAzurermRoleManagementPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/role_management_policy#timeouts DataAzurermRoleManagementPolicy#timeouts}
	Timeouts *DataAzurermRoleManagementPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

