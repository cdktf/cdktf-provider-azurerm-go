// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicapptriggerrecurrence


type LogicAppTriggerRecurrenceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/logic_app_trigger_recurrence#create LogicAppTriggerRecurrence#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/logic_app_trigger_recurrence#delete LogicAppTriggerRecurrence#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/logic_app_trigger_recurrence#read LogicAppTriggerRecurrence#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/logic_app_trigger_recurrence#update LogicAppTriggerRecurrence#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

