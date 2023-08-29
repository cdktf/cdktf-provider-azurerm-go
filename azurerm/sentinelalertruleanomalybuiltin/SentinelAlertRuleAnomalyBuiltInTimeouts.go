// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertruleanomalybuiltin


type SentinelAlertRuleAnomalyBuiltInTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_anomaly_built_in#create SentinelAlertRuleAnomalyBuiltIn#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_anomaly_built_in#delete SentinelAlertRuleAnomalyBuiltIn#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_anomaly_built_in#read SentinelAlertRuleAnomalyBuiltIn#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_anomaly_built_in#update SentinelAlertRuleAnomalyBuiltIn#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

