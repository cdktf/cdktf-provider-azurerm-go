// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policysetdefinition


type PolicySetDefinitionPolicyDefinitionGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/policy_set_definition#name PolicySetDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/policy_set_definition#additional_metadata_resource_id PolicySetDefinition#additional_metadata_resource_id}.
	AdditionalMetadataResourceId *string `field:"optional" json:"additionalMetadataResourceId" yaml:"additionalMetadataResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/policy_set_definition#category PolicySetDefinition#category}.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/policy_set_definition#description PolicySetDefinition#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/policy_set_definition#display_name PolicySetDefinition#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

