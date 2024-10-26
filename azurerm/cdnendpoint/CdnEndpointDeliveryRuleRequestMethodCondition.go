// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint


type CdnEndpointDeliveryRuleRequestMethodCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"required" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

