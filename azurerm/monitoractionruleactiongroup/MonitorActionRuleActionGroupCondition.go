// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractionruleactiongroup


type MonitorActionRuleActionGroupCondition struct {
	// alert_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_rule_action_group#alert_context MonitorActionRuleActionGroup#alert_context}
	AlertContext *MonitorActionRuleActionGroupConditionAlertContext `field:"optional" json:"alertContext" yaml:"alertContext"`
	// alert_rule_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_rule_action_group#alert_rule_id MonitorActionRuleActionGroup#alert_rule_id}
	AlertRuleId *MonitorActionRuleActionGroupConditionAlertRuleId `field:"optional" json:"alertRuleId" yaml:"alertRuleId"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_rule_action_group#description MonitorActionRuleActionGroup#description}
	Description *MonitorActionRuleActionGroupConditionDescription `field:"optional" json:"description" yaml:"description"`
	// monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_rule_action_group#monitor MonitorActionRuleActionGroup#monitor}
	Monitor *MonitorActionRuleActionGroupConditionMonitor `field:"optional" json:"monitor" yaml:"monitor"`
	// monitor_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_rule_action_group#monitor_service MonitorActionRuleActionGroup#monitor_service}
	MonitorService *MonitorActionRuleActionGroupConditionMonitorService `field:"optional" json:"monitorService" yaml:"monitorService"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_rule_action_group#severity MonitorActionRuleActionGroup#severity}
	Severity *MonitorActionRuleActionGroupConditionSeverity `field:"optional" json:"severity" yaml:"severity"`
	// target_resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_rule_action_group#target_resource_type MonitorActionRuleActionGroup#target_resource_type}
	TargetResourceType *MonitorActionRuleActionGroupConditionTargetResourceType `field:"optional" json:"targetResourceType" yaml:"targetResourceType"`
}

