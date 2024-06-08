// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayWafConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#enabled ApplicationGateway#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#firewall_mode ApplicationGateway#firewall_mode}.
	FirewallMode *string `field:"required" json:"firewallMode" yaml:"firewallMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#rule_set_version ApplicationGateway#rule_set_version}.
	RuleSetVersion *string `field:"required" json:"ruleSetVersion" yaml:"ruleSetVersion"`
	// disabled_rule_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#disabled_rule_group ApplicationGateway#disabled_rule_group}
	DisabledRuleGroup interface{} `field:"optional" json:"disabledRuleGroup" yaml:"disabledRuleGroup"`
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#exclusion ApplicationGateway#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#file_upload_limit_mb ApplicationGateway#file_upload_limit_mb}.
	FileUploadLimitMb *float64 `field:"optional" json:"fileUploadLimitMb" yaml:"fileUploadLimitMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#max_request_body_size_kb ApplicationGateway#max_request_body_size_kb}.
	MaxRequestBodySizeKb *float64 `field:"optional" json:"maxRequestBodySizeKb" yaml:"maxRequestBodySizeKb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#request_body_check ApplicationGateway#request_body_check}.
	RequestBodyCheck interface{} `field:"optional" json:"requestBodyCheck" yaml:"requestBodyCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/application_gateway#rule_set_type ApplicationGateway#rule_set_type}.
	RuleSetType *string `field:"optional" json:"ruleSetType" yaml:"ruleSetType"`
}

