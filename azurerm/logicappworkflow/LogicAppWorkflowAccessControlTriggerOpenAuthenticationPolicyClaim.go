// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappworkflow


type LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaim struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/logic_app_workflow#name LogicAppWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/logic_app_workflow#value LogicAppWorkflow#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

