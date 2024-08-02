// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionCondition struct {
	// alert_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#alert_context MonitorAlertProcessingRuleSuppression#alert_context}
	AlertContext *MonitorAlertProcessingRuleSuppressionConditionAlertContext `field:"optional" json:"alertContext" yaml:"alertContext"`
	// alert_rule_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#alert_rule_id MonitorAlertProcessingRuleSuppression#alert_rule_id}
	AlertRuleId *MonitorAlertProcessingRuleSuppressionConditionAlertRuleId `field:"optional" json:"alertRuleId" yaml:"alertRuleId"`
	// alert_rule_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#alert_rule_name MonitorAlertProcessingRuleSuppression#alert_rule_name}
	AlertRuleName *MonitorAlertProcessingRuleSuppressionConditionAlertRuleName `field:"optional" json:"alertRuleName" yaml:"alertRuleName"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#description MonitorAlertProcessingRuleSuppression#description}
	Description *MonitorAlertProcessingRuleSuppressionConditionDescription `field:"optional" json:"description" yaml:"description"`
	// monitor_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#monitor_condition MonitorAlertProcessingRuleSuppression#monitor_condition}
	MonitorCondition *MonitorAlertProcessingRuleSuppressionConditionMonitorCondition `field:"optional" json:"monitorCondition" yaml:"monitorCondition"`
	// monitor_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#monitor_service MonitorAlertProcessingRuleSuppression#monitor_service}
	MonitorService *MonitorAlertProcessingRuleSuppressionConditionMonitorService `field:"optional" json:"monitorService" yaml:"monitorService"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#severity MonitorAlertProcessingRuleSuppression#severity}
	Severity *MonitorAlertProcessingRuleSuppressionConditionSeverity `field:"optional" json:"severity" yaml:"severity"`
	// signal_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#signal_type MonitorAlertProcessingRuleSuppression#signal_type}
	SignalType *MonitorAlertProcessingRuleSuppressionConditionSignalType `field:"optional" json:"signalType" yaml:"signalType"`
	// target_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#target_resource MonitorAlertProcessingRuleSuppression#target_resource}
	TargetResource *MonitorAlertProcessingRuleSuppressionConditionTargetResource `field:"optional" json:"targetResource" yaml:"targetResource"`
	// target_resource_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#target_resource_group MonitorAlertProcessingRuleSuppression#target_resource_group}
	TargetResourceGroup *MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroup `field:"optional" json:"targetResourceGroup" yaml:"targetResourceGroup"`
	// target_resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_alert_processing_rule_suppression#target_resource_type MonitorAlertProcessingRuleSuppression#target_resource_type}
	TargetResourceType *MonitorAlertProcessingRuleSuppressionConditionTargetResourceType `field:"optional" json:"targetResourceType" yaml:"targetResourceType"`
}

