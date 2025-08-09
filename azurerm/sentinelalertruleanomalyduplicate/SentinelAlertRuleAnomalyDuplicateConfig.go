// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertruleanomalyduplicate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleAnomalyDuplicateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#built_in_rule_id SentinelAlertRuleAnomalyDuplicate#built_in_rule_id}.
	BuiltInRuleId *string `field:"required" json:"builtInRuleId" yaml:"builtInRuleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#display_name SentinelAlertRuleAnomalyDuplicate#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#enabled SentinelAlertRuleAnomalyDuplicate#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#log_analytics_workspace_id SentinelAlertRuleAnomalyDuplicate#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#mode SentinelAlertRuleAnomalyDuplicate#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#id SentinelAlertRuleAnomalyDuplicate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// multi_select_observation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#multi_select_observation SentinelAlertRuleAnomalyDuplicate#multi_select_observation}
	MultiSelectObservation interface{} `field:"optional" json:"multiSelectObservation" yaml:"multiSelectObservation"`
	// prioritized_exclude_observation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#prioritized_exclude_observation SentinelAlertRuleAnomalyDuplicate#prioritized_exclude_observation}
	PrioritizedExcludeObservation interface{} `field:"optional" json:"prioritizedExcludeObservation" yaml:"prioritizedExcludeObservation"`
	// single_select_observation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#single_select_observation SentinelAlertRuleAnomalyDuplicate#single_select_observation}
	SingleSelectObservation interface{} `field:"optional" json:"singleSelectObservation" yaml:"singleSelectObservation"`
	// threshold_observation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#threshold_observation SentinelAlertRuleAnomalyDuplicate#threshold_observation}
	ThresholdObservation interface{} `field:"optional" json:"thresholdObservation" yaml:"thresholdObservation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#timeouts SentinelAlertRuleAnomalyDuplicate#timeouts}
	Timeouts *SentinelAlertRuleAnomalyDuplicateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

