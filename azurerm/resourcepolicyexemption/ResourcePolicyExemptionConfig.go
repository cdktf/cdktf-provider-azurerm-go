// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcepolicyexemption

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourcePolicyExemptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#exemption_category ResourcePolicyExemption#exemption_category}.
	ExemptionCategory *string `field:"required" json:"exemptionCategory" yaml:"exemptionCategory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#name ResourcePolicyExemption#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#policy_assignment_id ResourcePolicyExemption#policy_assignment_id}.
	PolicyAssignmentId *string `field:"required" json:"policyAssignmentId" yaml:"policyAssignmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#resource_id ResourcePolicyExemption#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#description ResourcePolicyExemption#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#display_name ResourcePolicyExemption#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#expires_on ResourcePolicyExemption#expires_on}.
	ExpiresOn *string `field:"optional" json:"expiresOn" yaml:"expiresOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#id ResourcePolicyExemption#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#metadata ResourcePolicyExemption#metadata}.
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#policy_definition_reference_ids ResourcePolicyExemption#policy_definition_reference_ids}.
	PolicyDefinitionReferenceIds *[]*string `field:"optional" json:"policyDefinitionReferenceIds" yaml:"policyDefinitionReferenceIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/resource_policy_exemption#timeouts ResourcePolicyExemption#timeouts}
	Timeouts *ResourcePolicyExemptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

