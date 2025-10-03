// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadogmonitortagrule


type DatadogMonitorTagRuleMetricFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/datadog_monitor_tag_rule#action DatadogMonitorTagRule#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/datadog_monitor_tag_rule#name DatadogMonitorTagRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/datadog_monitor_tag_rule#value DatadogMonitorTagRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

