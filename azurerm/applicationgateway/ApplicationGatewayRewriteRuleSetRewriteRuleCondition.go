// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayRewriteRuleSetRewriteRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/application_gateway#pattern ApplicationGateway#pattern}.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/application_gateway#variable ApplicationGateway#variable}.
	Variable *string `field:"required" json:"variable" yaml:"variable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/application_gateway#ignore_case ApplicationGateway#ignore_case}.
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/application_gateway#negate ApplicationGateway#negate}.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
}

