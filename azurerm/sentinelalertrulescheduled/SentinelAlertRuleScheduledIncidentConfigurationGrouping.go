// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled


type SentinelAlertRuleScheduledIncidentConfigurationGrouping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/sentinel_alert_rule_scheduled#enabled SentinelAlertRuleScheduled#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/sentinel_alert_rule_scheduled#entity_matching_method SentinelAlertRuleScheduled#entity_matching_method}.
	EntityMatchingMethod *string `field:"optional" json:"entityMatchingMethod" yaml:"entityMatchingMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/sentinel_alert_rule_scheduled#group_by_alert_details SentinelAlertRuleScheduled#group_by_alert_details}.
	GroupByAlertDetails *[]*string `field:"optional" json:"groupByAlertDetails" yaml:"groupByAlertDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/sentinel_alert_rule_scheduled#group_by_custom_details SentinelAlertRuleScheduled#group_by_custom_details}.
	GroupByCustomDetails *[]*string `field:"optional" json:"groupByCustomDetails" yaml:"groupByCustomDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/sentinel_alert_rule_scheduled#group_by_entities SentinelAlertRuleScheduled#group_by_entities}.
	GroupByEntities *[]*string `field:"optional" json:"groupByEntities" yaml:"groupByEntities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/sentinel_alert_rule_scheduled#lookback_duration SentinelAlertRuleScheduled#lookback_duration}.
	LookbackDuration *string `field:"optional" json:"lookbackDuration" yaml:"lookbackDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/sentinel_alert_rule_scheduled#reopen_closed_incidents SentinelAlertRuleScheduled#reopen_closed_incidents}.
	ReopenClosedIncidents interface{} `field:"optional" json:"reopenClosedIncidents" yaml:"reopenClosedIncidents"`
}

