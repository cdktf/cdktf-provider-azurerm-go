// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupCondition struct {
	// alert_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#alert_context MonitorAlertProcessingRuleActionGroup#alert_context}
	AlertContext *MonitorAlertProcessingRuleActionGroupConditionAlertContext `field:"optional" json:"alertContext" yaml:"alertContext"`
	// alert_rule_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#alert_rule_id MonitorAlertProcessingRuleActionGroup#alert_rule_id}
	AlertRuleId *MonitorAlertProcessingRuleActionGroupConditionAlertRuleId `field:"optional" json:"alertRuleId" yaml:"alertRuleId"`
	// alert_rule_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#alert_rule_name MonitorAlertProcessingRuleActionGroup#alert_rule_name}
	AlertRuleName *MonitorAlertProcessingRuleActionGroupConditionAlertRuleName `field:"optional" json:"alertRuleName" yaml:"alertRuleName"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#description MonitorAlertProcessingRuleActionGroup#description}
	Description *MonitorAlertProcessingRuleActionGroupConditionDescription `field:"optional" json:"description" yaml:"description"`
	// monitor_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#monitor_condition MonitorAlertProcessingRuleActionGroup#monitor_condition}
	MonitorCondition *MonitorAlertProcessingRuleActionGroupConditionMonitorCondition `field:"optional" json:"monitorCondition" yaml:"monitorCondition"`
	// monitor_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#monitor_service MonitorAlertProcessingRuleActionGroup#monitor_service}
	MonitorService *MonitorAlertProcessingRuleActionGroupConditionMonitorService `field:"optional" json:"monitorService" yaml:"monitorService"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#severity MonitorAlertProcessingRuleActionGroup#severity}
	Severity *MonitorAlertProcessingRuleActionGroupConditionSeverity `field:"optional" json:"severity" yaml:"severity"`
	// signal_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#signal_type MonitorAlertProcessingRuleActionGroup#signal_type}
	SignalType *MonitorAlertProcessingRuleActionGroupConditionSignalType `field:"optional" json:"signalType" yaml:"signalType"`
	// target_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#target_resource MonitorAlertProcessingRuleActionGroup#target_resource}
	TargetResource *MonitorAlertProcessingRuleActionGroupConditionTargetResource `field:"optional" json:"targetResource" yaml:"targetResource"`
	// target_resource_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#target_resource_group MonitorAlertProcessingRuleActionGroup#target_resource_group}
	TargetResourceGroup *MonitorAlertProcessingRuleActionGroupConditionTargetResourceGroup `field:"optional" json:"targetResourceGroup" yaml:"targetResourceGroup"`
	// target_resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_alert_processing_rule_action_group#target_resource_type MonitorAlertProcessingRuleActionGroup#target_resource_type}
	TargetResourceType *MonitorAlertProcessingRuleActionGroupConditionTargetResourceType `field:"optional" json:"targetResourceType" yaml:"targetResourceType"`
}

