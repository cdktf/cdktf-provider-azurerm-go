// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappworkflow


type LogicAppWorkflowAccessControl struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/logic_app_workflow#action LogicAppWorkflow#action}
	Action *LogicAppWorkflowAccessControlAction `field:"optional" json:"action" yaml:"action"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/logic_app_workflow#content LogicAppWorkflow#content}
	Content *LogicAppWorkflowAccessControlContent `field:"optional" json:"content" yaml:"content"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/logic_app_workflow#trigger LogicAppWorkflow#trigger}
	Trigger *LogicAppWorkflowAccessControlTrigger `field:"optional" json:"trigger" yaml:"trigger"`
	// workflow_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/logic_app_workflow#workflow_management LogicAppWorkflow#workflow_management}
	WorkflowManagement *LogicAppWorkflowAccessControlWorkflowManagement `field:"optional" json:"workflowManagement" yaml:"workflowManagement"`
}

