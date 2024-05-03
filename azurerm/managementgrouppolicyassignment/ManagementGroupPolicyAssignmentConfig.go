// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicyassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagementGroupPolicyAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#management_group_id ManagementGroupPolicyAssignment#management_group_id}.
	ManagementGroupId *string `field:"required" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#name ManagementGroupPolicyAssignment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#policy_definition_id ManagementGroupPolicyAssignment#policy_definition_id}.
	PolicyDefinitionId *string `field:"required" json:"policyDefinitionId" yaml:"policyDefinitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#description ManagementGroupPolicyAssignment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#display_name ManagementGroupPolicyAssignment#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#enforce ManagementGroupPolicyAssignment#enforce}.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#id ManagementGroupPolicyAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#identity ManagementGroupPolicyAssignment#identity}
	Identity *ManagementGroupPolicyAssignmentIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#location ManagementGroupPolicyAssignment#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#metadata ManagementGroupPolicyAssignment#metadata}.
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// non_compliance_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#non_compliance_message ManagementGroupPolicyAssignment#non_compliance_message}
	NonComplianceMessage interface{} `field:"optional" json:"nonComplianceMessage" yaml:"nonComplianceMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#not_scopes ManagementGroupPolicyAssignment#not_scopes}.
	NotScopes *[]*string `field:"optional" json:"notScopes" yaml:"notScopes"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#overrides ManagementGroupPolicyAssignment#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#parameters ManagementGroupPolicyAssignment#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// resource_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#resource_selectors ManagementGroupPolicyAssignment#resource_selectors}
	ResourceSelectors interface{} `field:"optional" json:"resourceSelectors" yaml:"resourceSelectors"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#timeouts ManagementGroupPolicyAssignment#timeouts}
	Timeouts *ManagementGroupPolicyAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

