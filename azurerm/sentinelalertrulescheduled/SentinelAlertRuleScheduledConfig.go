// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleScheduledConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#display_name SentinelAlertRuleScheduled#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#log_analytics_workspace_id SentinelAlertRuleScheduled#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#name SentinelAlertRuleScheduled#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#query SentinelAlertRuleScheduled#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#severity SentinelAlertRuleScheduled#severity}.
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// alert_details_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#alert_details_override SentinelAlertRuleScheduled#alert_details_override}
	AlertDetailsOverride interface{} `field:"optional" json:"alertDetailsOverride" yaml:"alertDetailsOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#alert_rule_template_guid SentinelAlertRuleScheduled#alert_rule_template_guid}.
	AlertRuleTemplateGuid *string `field:"optional" json:"alertRuleTemplateGuid" yaml:"alertRuleTemplateGuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#alert_rule_template_version SentinelAlertRuleScheduled#alert_rule_template_version}.
	AlertRuleTemplateVersion *string `field:"optional" json:"alertRuleTemplateVersion" yaml:"alertRuleTemplateVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#custom_details SentinelAlertRuleScheduled#custom_details}.
	CustomDetails *map[string]*string `field:"optional" json:"customDetails" yaml:"customDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#description SentinelAlertRuleScheduled#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#enabled SentinelAlertRuleScheduled#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// entity_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#entity_mapping SentinelAlertRuleScheduled#entity_mapping}
	EntityMapping interface{} `field:"optional" json:"entityMapping" yaml:"entityMapping"`
	// event_grouping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#event_grouping SentinelAlertRuleScheduled#event_grouping}
	EventGrouping *SentinelAlertRuleScheduledEventGrouping `field:"optional" json:"eventGrouping" yaml:"eventGrouping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#id SentinelAlertRuleScheduled#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// incident block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#incident SentinelAlertRuleScheduled#incident}
	Incident *SentinelAlertRuleScheduledIncident `field:"optional" json:"incident" yaml:"incident"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#query_frequency SentinelAlertRuleScheduled#query_frequency}.
	QueryFrequency *string `field:"optional" json:"queryFrequency" yaml:"queryFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#query_period SentinelAlertRuleScheduled#query_period}.
	QueryPeriod *string `field:"optional" json:"queryPeriod" yaml:"queryPeriod"`
	// sentinel_entity_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#sentinel_entity_mapping SentinelAlertRuleScheduled#sentinel_entity_mapping}
	SentinelEntityMapping interface{} `field:"optional" json:"sentinelEntityMapping" yaml:"sentinelEntityMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#suppression_duration SentinelAlertRuleScheduled#suppression_duration}.
	SuppressionDuration *string `field:"optional" json:"suppressionDuration" yaml:"suppressionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#suppression_enabled SentinelAlertRuleScheduled#suppression_enabled}.
	SuppressionEnabled interface{} `field:"optional" json:"suppressionEnabled" yaml:"suppressionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#tactics SentinelAlertRuleScheduled#tactics}.
	Tactics *[]*string `field:"optional" json:"tactics" yaml:"tactics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#techniques SentinelAlertRuleScheduled#techniques}.
	Techniques *[]*string `field:"optional" json:"techniques" yaml:"techniques"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#timeouts SentinelAlertRuleScheduled#timeouts}
	Timeouts *SentinelAlertRuleScheduledTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#trigger_operator SentinelAlertRuleScheduled#trigger_operator}.
	TriggerOperator *string `field:"optional" json:"triggerOperator" yaml:"triggerOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/sentinel_alert_rule_scheduled#trigger_threshold SentinelAlertRuleScheduled#trigger_threshold}.
	TriggerThreshold *float64 `field:"optional" json:"triggerThreshold" yaml:"triggerThreshold"`
}

