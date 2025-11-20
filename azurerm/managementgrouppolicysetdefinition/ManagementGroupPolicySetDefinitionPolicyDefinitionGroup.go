// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicysetdefinition


type ManagementGroupPolicySetDefinitionPolicyDefinitionGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/management_group_policy_set_definition#name ManagementGroupPolicySetDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/management_group_policy_set_definition#additional_metadata_resource_id ManagementGroupPolicySetDefinition#additional_metadata_resource_id}.
	AdditionalMetadataResourceId *string `field:"optional" json:"additionalMetadataResourceId" yaml:"additionalMetadataResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/management_group_policy_set_definition#category ManagementGroupPolicySetDefinition#category}.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/management_group_policy_set_definition#description ManagementGroupPolicySetDefinition#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/management_group_policy_set_definition#display_name ManagementGroupPolicySetDefinition#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

