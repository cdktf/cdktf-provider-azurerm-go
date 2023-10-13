// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcepolicyassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourcePolicyAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#name ResourcePolicyAssignment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#policy_definition_id ResourcePolicyAssignment#policy_definition_id}.
	PolicyDefinitionId *string `field:"required" json:"policyDefinitionId" yaml:"policyDefinitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#resource_id ResourcePolicyAssignment#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#description ResourcePolicyAssignment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#display_name ResourcePolicyAssignment#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#enforce ResourcePolicyAssignment#enforce}.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#id ResourcePolicyAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#identity ResourcePolicyAssignment#identity}
	Identity *ResourcePolicyAssignmentIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#location ResourcePolicyAssignment#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#metadata ResourcePolicyAssignment#metadata}.
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// non_compliance_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#non_compliance_message ResourcePolicyAssignment#non_compliance_message}
	NonComplianceMessage interface{} `field:"optional" json:"nonComplianceMessage" yaml:"nonComplianceMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#not_scopes ResourcePolicyAssignment#not_scopes}.
	NotScopes *[]*string `field:"optional" json:"notScopes" yaml:"notScopes"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#overrides ResourcePolicyAssignment#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#parameters ResourcePolicyAssignment#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// resource_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#resource_selectors ResourcePolicyAssignment#resource_selectors}
	ResourceSelectors interface{} `field:"optional" json:"resourceSelectors" yaml:"resourceSelectors"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_policy_assignment#timeouts ResourcePolicyAssignment#timeouts}
	Timeouts *ResourcePolicyAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

