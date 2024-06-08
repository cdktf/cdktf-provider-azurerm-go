// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractionrulesuppression


type MonitorActionRuleSuppressionSuppression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/monitor_action_rule_suppression#recurrence_type MonitorActionRuleSuppression#recurrence_type}.
	RecurrenceType *string `field:"required" json:"recurrenceType" yaml:"recurrenceType"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/monitor_action_rule_suppression#schedule MonitorActionRuleSuppression#schedule}
	Schedule *MonitorActionRuleSuppressionSuppressionSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

