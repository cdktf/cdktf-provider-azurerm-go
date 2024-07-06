// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled


type SentinelAlertRuleScheduledIncidentGrouping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_alert_rule_scheduled#by_alert_details SentinelAlertRuleScheduled#by_alert_details}.
	ByAlertDetails *[]*string `field:"optional" json:"byAlertDetails" yaml:"byAlertDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_alert_rule_scheduled#by_custom_details SentinelAlertRuleScheduled#by_custom_details}.
	ByCustomDetails *[]*string `field:"optional" json:"byCustomDetails" yaml:"byCustomDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_alert_rule_scheduled#by_entities SentinelAlertRuleScheduled#by_entities}.
	ByEntities *[]*string `field:"optional" json:"byEntities" yaml:"byEntities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_alert_rule_scheduled#enabled SentinelAlertRuleScheduled#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_alert_rule_scheduled#entity_matching_method SentinelAlertRuleScheduled#entity_matching_method}.
	EntityMatchingMethod *string `field:"optional" json:"entityMatchingMethod" yaml:"entityMatchingMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_alert_rule_scheduled#lookback_duration SentinelAlertRuleScheduled#lookback_duration}.
	LookbackDuration *string `field:"optional" json:"lookbackDuration" yaml:"lookbackDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_alert_rule_scheduled#reopen_closed_incidents SentinelAlertRuleScheduled#reopen_closed_incidents}.
	ReopenClosedIncidents interface{} `field:"optional" json:"reopenClosedIncidents" yaml:"reopenClosedIncidents"`
}

