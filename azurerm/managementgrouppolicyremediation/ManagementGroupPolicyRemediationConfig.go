// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicyremediation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagementGroupPolicyRemediationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#management_group_id ManagementGroupPolicyRemediation#management_group_id}.
	ManagementGroupId *string `field:"required" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#name ManagementGroupPolicyRemediation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#policy_assignment_id ManagementGroupPolicyRemediation#policy_assignment_id}.
	PolicyAssignmentId *string `field:"required" json:"policyAssignmentId" yaml:"policyAssignmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#failure_percentage ManagementGroupPolicyRemediation#failure_percentage}.
	FailurePercentage *float64 `field:"optional" json:"failurePercentage" yaml:"failurePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#id ManagementGroupPolicyRemediation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#location_filters ManagementGroupPolicyRemediation#location_filters}.
	LocationFilters *[]*string `field:"optional" json:"locationFilters" yaml:"locationFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#parallel_deployments ManagementGroupPolicyRemediation#parallel_deployments}.
	ParallelDeployments *float64 `field:"optional" json:"parallelDeployments" yaml:"parallelDeployments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#policy_definition_id ManagementGroupPolicyRemediation#policy_definition_id}.
	PolicyDefinitionId *string `field:"optional" json:"policyDefinitionId" yaml:"policyDefinitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#policy_definition_reference_id ManagementGroupPolicyRemediation#policy_definition_reference_id}.
	PolicyDefinitionReferenceId *string `field:"optional" json:"policyDefinitionReferenceId" yaml:"policyDefinitionReferenceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#resource_count ManagementGroupPolicyRemediation#resource_count}.
	ResourceCount *float64 `field:"optional" json:"resourceCount" yaml:"resourceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#resource_discovery_mode ManagementGroupPolicyRemediation#resource_discovery_mode}.
	ResourceDiscoveryMode *string `field:"optional" json:"resourceDiscoveryMode" yaml:"resourceDiscoveryMode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/management_group_policy_remediation#timeouts ManagementGroupPolicyRemediation#timeouts}
	Timeouts *ManagementGroupPolicyRemediationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

