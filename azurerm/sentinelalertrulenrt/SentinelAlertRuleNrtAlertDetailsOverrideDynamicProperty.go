// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulenrt


type SentinelAlertRuleNrtAlertDetailsOverrideDynamicProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_nrt#name SentinelAlertRuleNrt#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/sentinel_alert_rule_nrt#value SentinelAlertRuleNrt#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

