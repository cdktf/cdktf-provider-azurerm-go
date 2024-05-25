// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertruleanomalyduplicate


type SentinelAlertRuleAnomalyDuplicateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#create SentinelAlertRuleAnomalyDuplicate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#delete SentinelAlertRuleAnomalyDuplicate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#read SentinelAlertRuleAnomalyDuplicate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sentinel_alert_rule_anomaly_duplicate#update SentinelAlertRuleAnomalyDuplicate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

