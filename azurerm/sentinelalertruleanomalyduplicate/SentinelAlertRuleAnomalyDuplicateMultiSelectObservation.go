// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertruleanomalyduplicate


type SentinelAlertRuleAnomalyDuplicateMultiSelectObservation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#name SentinelAlertRuleAnomalyDuplicate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#values SentinelAlertRuleAnomalyDuplicate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

