// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertruleanomalyduplicate


type SentinelAlertRuleAnomalyDuplicatePrioritizedExcludeObservation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#name SentinelAlertRuleAnomalyDuplicate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#exclude SentinelAlertRuleAnomalyDuplicate#exclude}.
	Exclude *string `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#prioritize SentinelAlertRuleAnomalyDuplicate#prioritize}.
	Prioritize *string `field:"optional" json:"prioritize" yaml:"prioritize"`
}

