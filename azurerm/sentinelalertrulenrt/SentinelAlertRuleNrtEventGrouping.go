// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulenrt


type SentinelAlertRuleNrtEventGrouping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/sentinel_alert_rule_nrt#aggregation_method SentinelAlertRuleNrt#aggregation_method}.
	AggregationMethod *string `field:"required" json:"aggregationMethod" yaml:"aggregationMethod"`
}

