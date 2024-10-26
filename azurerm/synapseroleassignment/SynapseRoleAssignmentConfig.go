// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseroleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SynapseRoleAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/synapse_role_assignment#principal_id SynapseRoleAssignment#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/synapse_role_assignment#role_name SynapseRoleAssignment#role_name}.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/synapse_role_assignment#id SynapseRoleAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/synapse_role_assignment#principal_type SynapseRoleAssignment#principal_type}.
	PrincipalType *string `field:"optional" json:"principalType" yaml:"principalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/synapse_role_assignment#synapse_spark_pool_id SynapseRoleAssignment#synapse_spark_pool_id}.
	SynapseSparkPoolId *string `field:"optional" json:"synapseSparkPoolId" yaml:"synapseSparkPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/synapse_role_assignment#synapse_workspace_id SynapseRoleAssignment#synapse_workspace_id}.
	SynapseWorkspaceId *string `field:"optional" json:"synapseWorkspaceId" yaml:"synapseWorkspaceId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/synapse_role_assignment#timeouts SynapseRoleAssignment#timeouts}
	Timeouts *SynapseRoleAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

