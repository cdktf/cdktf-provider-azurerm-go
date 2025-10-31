// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulefusion


type SentinelAlertRuleFusionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_alert_rule_fusion#create SentinelAlertRuleFusion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_alert_rule_fusion#delete SentinelAlertRuleFusion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_alert_rule_fusion#read SentinelAlertRuleFusion#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_alert_rule_fusion#update SentinelAlertRuleFusion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

