// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policysetdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicySetDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#display_name PolicySetDefinition#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#name PolicySetDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// policy_definition_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#policy_definition_reference PolicySetDefinition#policy_definition_reference}
	PolicyDefinitionReference interface{} `field:"required" json:"policyDefinitionReference" yaml:"policyDefinitionReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#policy_type PolicySetDefinition#policy_type}.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#description PolicySetDefinition#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#id PolicySetDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#management_group_id PolicySetDefinition#management_group_id}.
	ManagementGroupId *string `field:"optional" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#metadata PolicySetDefinition#metadata}.
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#parameters PolicySetDefinition#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// policy_definition_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#policy_definition_group PolicySetDefinition#policy_definition_group}
	PolicyDefinitionGroup interface{} `field:"optional" json:"policyDefinitionGroup" yaml:"policyDefinitionGroup"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/policy_set_definition#timeouts PolicySetDefinition#timeouts}
	Timeouts *PolicySetDefinitionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

