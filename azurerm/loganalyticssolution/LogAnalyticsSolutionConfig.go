// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticssolution

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogAnalyticsSolutionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#location LogAnalyticsSolution#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#plan LogAnalyticsSolution#plan}
	Plan *LogAnalyticsSolutionPlan `field:"required" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#resource_group_name LogAnalyticsSolution#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#solution_name LogAnalyticsSolution#solution_name}.
	SolutionName *string `field:"required" json:"solutionName" yaml:"solutionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#workspace_name LogAnalyticsSolution#workspace_name}.
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#workspace_resource_id LogAnalyticsSolution#workspace_resource_id}.
	WorkspaceResourceId *string `field:"required" json:"workspaceResourceId" yaml:"workspaceResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#id LogAnalyticsSolution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#tags LogAnalyticsSolution#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/log_analytics_solution#timeouts LogAnalyticsSolution#timeouts}
	Timeouts *LogAnalyticsSolutionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

