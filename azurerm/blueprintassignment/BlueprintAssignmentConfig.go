// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package blueprintassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BlueprintAssignmentConfig struct {
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
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#identity BlueprintAssignment#identity}
	Identity *BlueprintAssignmentIdentity `field:"required" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#location BlueprintAssignment#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#name BlueprintAssignment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#target_subscription_id BlueprintAssignment#target_subscription_id}.
	TargetSubscriptionId *string `field:"required" json:"targetSubscriptionId" yaml:"targetSubscriptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#version_id BlueprintAssignment#version_id}.
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#id BlueprintAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#lock_exclude_actions BlueprintAssignment#lock_exclude_actions}.
	LockExcludeActions *[]*string `field:"optional" json:"lockExcludeActions" yaml:"lockExcludeActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#lock_exclude_principals BlueprintAssignment#lock_exclude_principals}.
	LockExcludePrincipals *[]*string `field:"optional" json:"lockExcludePrincipals" yaml:"lockExcludePrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#lock_mode BlueprintAssignment#lock_mode}.
	LockMode *string `field:"optional" json:"lockMode" yaml:"lockMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#parameter_values BlueprintAssignment#parameter_values}.
	ParameterValues *string `field:"optional" json:"parameterValues" yaml:"parameterValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#resource_groups BlueprintAssignment#resource_groups}.
	ResourceGroups *string `field:"optional" json:"resourceGroups" yaml:"resourceGroups"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/blueprint_assignment#timeouts BlueprintAssignment#timeouts}
	Timeouts *BlueprintAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

