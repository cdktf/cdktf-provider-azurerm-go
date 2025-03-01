// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatracetagrules


type DynatraceTagRulesMetricRule struct {
	// filtering_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/dynatrace_tag_rules#filtering_tag DynatraceTagRules#filtering_tag}
	FilteringTag interface{} `field:"required" json:"filteringTag" yaml:"filteringTag"`
}

