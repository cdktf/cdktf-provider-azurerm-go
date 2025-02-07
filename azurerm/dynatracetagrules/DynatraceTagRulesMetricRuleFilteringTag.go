// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatracetagrules


type DynatraceTagRulesMetricRuleFilteringTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/dynatrace_tag_rules#action DynatraceTagRules#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/dynatrace_tag_rules#name DynatraceTagRules#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/dynatrace_tag_rules#value DynatraceTagRules#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

