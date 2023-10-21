// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logzsubaccounttagrule


type LogzSubAccountTagRuleTagFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/logz_sub_account_tag_rule#action LogzSubAccountTagRule#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/logz_sub_account_tag_rule#name LogzSubAccountTagRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/logz_sub_account_tag_rule#value LogzSubAccountTagRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

