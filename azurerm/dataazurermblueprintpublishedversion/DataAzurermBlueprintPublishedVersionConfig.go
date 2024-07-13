// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermblueprintpublishedversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermBlueprintPublishedVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/data-sources/blueprint_published_version#blueprint_name DataAzurermBlueprintPublishedVersion#blueprint_name}.
	BlueprintName *string `field:"required" json:"blueprintName" yaml:"blueprintName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/data-sources/blueprint_published_version#scope_id DataAzurermBlueprintPublishedVersion#scope_id}.
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/data-sources/blueprint_published_version#version DataAzurermBlueprintPublishedVersion#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/data-sources/blueprint_published_version#id DataAzurermBlueprintPublishedVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/data-sources/blueprint_published_version#timeouts DataAzurermBlueprintPublishedVersion#timeouts}
	Timeouts *DataAzurermBlueprintPublishedVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

