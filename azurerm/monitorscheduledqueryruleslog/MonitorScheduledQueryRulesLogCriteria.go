// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorscheduledqueryruleslog


type MonitorScheduledQueryRulesLogCriteria struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_scheduled_query_rules_log#dimension MonitorScheduledQueryRulesLog#dimension}
	Dimension interface{} `field:"required" json:"dimension" yaml:"dimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_scheduled_query_rules_log#metric_name MonitorScheduledQueryRulesLog#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
}

