// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorscheduledqueryrulesalert


type MonitorScheduledQueryRulesAlertTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/monitor_scheduled_query_rules_alert#operator MonitorScheduledQueryRulesAlert#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/monitor_scheduled_query_rules_alert#threshold MonitorScheduledQueryRulesAlert#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// metric_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/monitor_scheduled_query_rules_alert#metric_trigger MonitorScheduledQueryRulesAlert#metric_trigger}
	MetricTrigger *MonitorScheduledQueryRulesAlertTriggerMetricTrigger `field:"optional" json:"metricTrigger" yaml:"metricTrigger"`
}

