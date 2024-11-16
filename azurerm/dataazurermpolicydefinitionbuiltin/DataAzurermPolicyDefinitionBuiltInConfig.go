// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermpolicydefinitionbuiltin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermPolicyDefinitionBuiltInConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/data-sources/policy_definition_built_in#display_name DataAzurermPolicyDefinitionBuiltIn#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/data-sources/policy_definition_built_in#id DataAzurermPolicyDefinitionBuiltIn#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/data-sources/policy_definition_built_in#management_group_name DataAzurermPolicyDefinitionBuiltIn#management_group_name}.
	ManagementGroupName *string `field:"optional" json:"managementGroupName" yaml:"managementGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/data-sources/policy_definition_built_in#name DataAzurermPolicyDefinitionBuiltIn#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/data-sources/policy_definition_built_in#timeouts DataAzurermPolicyDefinitionBuiltIn#timeouts}
	Timeouts *DataAzurermPolicyDefinitionBuiltInTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

