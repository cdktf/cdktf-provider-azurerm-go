// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorscheduledqueryruleslog


type MonitorScheduledQueryRulesLogCriteriaDimension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_scheduled_query_rules_log#name MonitorScheduledQueryRulesLog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_scheduled_query_rules_log#values MonitorScheduledQueryRulesLog#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_scheduled_query_rules_log#operator MonitorScheduledQueryRulesLog#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

