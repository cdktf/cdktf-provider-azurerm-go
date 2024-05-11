// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorscheduledqueryrulesalertv2


type MonitorScheduledQueryRulesAlertV2Identity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/monitor_scheduled_query_rules_alert_v2#type MonitorScheduledQueryRulesAlertV2#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/monitor_scheduled_query_rules_alert_v2#identity_ids MonitorScheduledQueryRulesAlertV2#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

