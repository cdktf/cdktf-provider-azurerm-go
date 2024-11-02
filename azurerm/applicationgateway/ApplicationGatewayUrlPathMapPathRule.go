// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayUrlPathMapPathRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/application_gateway#paths ApplicationGateway#paths}.
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/application_gateway#backend_address_pool_name ApplicationGateway#backend_address_pool_name}.
	BackendAddressPoolName *string `field:"optional" json:"backendAddressPoolName" yaml:"backendAddressPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/application_gateway#backend_http_settings_name ApplicationGateway#backend_http_settings_name}.
	BackendHttpSettingsName *string `field:"optional" json:"backendHttpSettingsName" yaml:"backendHttpSettingsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/application_gateway#firewall_policy_id ApplicationGateway#firewall_policy_id}.
	FirewallPolicyId *string `field:"optional" json:"firewallPolicyId" yaml:"firewallPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/application_gateway#redirect_configuration_name ApplicationGateway#redirect_configuration_name}.
	RedirectConfigurationName *string `field:"optional" json:"redirectConfigurationName" yaml:"redirectConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/application_gateway#rewrite_rule_set_name ApplicationGateway#rewrite_rule_set_name}.
	RewriteRuleSetName *string `field:"optional" json:"rewriteRuleSetName" yaml:"rewriteRuleSetName"`
}

