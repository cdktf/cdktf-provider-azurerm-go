// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorscheduledqueryrulesalertv2


type MonitorScheduledQueryRulesAlertV2Action struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/monitor_scheduled_query_rules_alert_v2#action_groups MonitorScheduledQueryRulesAlertV2#action_groups}.
	ActionGroups *[]*string `field:"optional" json:"actionGroups" yaml:"actionGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/monitor_scheduled_query_rules_alert_v2#custom_properties MonitorScheduledQueryRulesAlertV2#custom_properties}.
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
}

