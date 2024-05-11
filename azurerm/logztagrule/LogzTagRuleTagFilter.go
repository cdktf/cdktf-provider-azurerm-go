// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logztagrule


type LogzTagRuleTagFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/logz_tag_rule#action LogzTagRule#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/logz_tag_rule#name LogzTagRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/logz_tag_rule#value LogzTagRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

