// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayRewriteRuleSetRewriteRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_gateway#rule_sequence ApplicationGateway#rule_sequence}.
	RuleSequence *float64 `field:"required" json:"ruleSequence" yaml:"ruleSequence"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_gateway#condition ApplicationGateway#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// request_header_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_gateway#request_header_configuration ApplicationGateway#request_header_configuration}
	RequestHeaderConfiguration interface{} `field:"optional" json:"requestHeaderConfiguration" yaml:"requestHeaderConfiguration"`
	// response_header_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_gateway#response_header_configuration ApplicationGateway#response_header_configuration}
	ResponseHeaderConfiguration interface{} `field:"optional" json:"responseHeaderConfiguration" yaml:"responseHeaderConfiguration"`
	// url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_gateway#url ApplicationGateway#url}
	Url *ApplicationGatewayRewriteRuleSetRewriteRuleUrl `field:"optional" json:"url" yaml:"url"`
}

