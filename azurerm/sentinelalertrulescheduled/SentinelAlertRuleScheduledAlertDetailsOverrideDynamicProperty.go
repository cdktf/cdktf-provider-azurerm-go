// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled


type SentinelAlertRuleScheduledAlertDetailsOverrideDynamicProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/sentinel_alert_rule_scheduled#name SentinelAlertRuleScheduled#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/sentinel_alert_rule_scheduled#value SentinelAlertRuleScheduled#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

