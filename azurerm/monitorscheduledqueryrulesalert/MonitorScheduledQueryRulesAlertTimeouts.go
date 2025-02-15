// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorscheduledqueryrulesalert


type MonitorScheduledQueryRulesAlertTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_scheduled_query_rules_alert#create MonitorScheduledQueryRulesAlert#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_scheduled_query_rules_alert#delete MonitorScheduledQueryRulesAlert#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_scheduled_query_rules_alert#read MonitorScheduledQueryRulesAlert#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_scheduled_query_rules_alert#update MonitorScheduledQueryRulesAlert#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

