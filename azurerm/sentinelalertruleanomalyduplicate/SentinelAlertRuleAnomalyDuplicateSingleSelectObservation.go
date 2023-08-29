// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertruleanomalyduplicate


type SentinelAlertRuleAnomalyDuplicateSingleSelectObservation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#name SentinelAlertRuleAnomalyDuplicate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#value SentinelAlertRuleAnomalyDuplicate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

