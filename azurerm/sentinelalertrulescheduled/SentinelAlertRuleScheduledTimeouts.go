// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled


type SentinelAlertRuleScheduledTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_scheduled#create SentinelAlertRuleScheduled#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_scheduled#delete SentinelAlertRuleScheduled#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_scheduled#read SentinelAlertRuleScheduled#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sentinel_alert_rule_scheduled#update SentinelAlertRuleScheduled#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

