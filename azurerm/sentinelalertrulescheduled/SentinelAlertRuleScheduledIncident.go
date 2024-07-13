// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled


type SentinelAlertRuleScheduledIncident struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/sentinel_alert_rule_scheduled#create_incident_enabled SentinelAlertRuleScheduled#create_incident_enabled}.
	CreateIncidentEnabled interface{} `field:"required" json:"createIncidentEnabled" yaml:"createIncidentEnabled"`
	// grouping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/sentinel_alert_rule_scheduled#grouping SentinelAlertRuleScheduled#grouping}
	Grouping *SentinelAlertRuleScheduledIncidentGrouping `field:"required" json:"grouping" yaml:"grouping"`
}

