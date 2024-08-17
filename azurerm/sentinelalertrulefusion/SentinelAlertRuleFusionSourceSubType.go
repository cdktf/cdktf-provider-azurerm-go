// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulefusion


type SentinelAlertRuleFusionSourceSubType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/sentinel_alert_rule_fusion#name SentinelAlertRuleFusion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/sentinel_alert_rule_fusion#severities_allowed SentinelAlertRuleFusion#severities_allowed}.
	SeveritiesAllowed *[]*string `field:"required" json:"severitiesAllowed" yaml:"severitiesAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/sentinel_alert_rule_fusion#enabled SentinelAlertRuleFusion#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

