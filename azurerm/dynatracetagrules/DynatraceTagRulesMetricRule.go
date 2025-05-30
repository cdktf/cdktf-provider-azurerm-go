// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatracetagrules


type DynatraceTagRulesMetricRule struct {
	// filtering_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/dynatrace_tag_rules#filtering_tag DynatraceTagRules#filtering_tag}
	FilteringTag interface{} `field:"required" json:"filteringTag" yaml:"filteringTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/dynatrace_tag_rules#sending_metrics_enabled DynatraceTagRules#sending_metrics_enabled}.
	SendingMetricsEnabled interface{} `field:"optional" json:"sendingMetricsEnabled" yaml:"sendingMetricsEnabled"`
}

