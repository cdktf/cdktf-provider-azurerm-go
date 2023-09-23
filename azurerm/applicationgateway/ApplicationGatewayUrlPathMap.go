// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayUrlPathMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// path_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/application_gateway#path_rule ApplicationGateway#path_rule}
	PathRule interface{} `field:"required" json:"pathRule" yaml:"pathRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/application_gateway#default_backend_address_pool_name ApplicationGateway#default_backend_address_pool_name}.
	DefaultBackendAddressPoolName *string `field:"optional" json:"defaultBackendAddressPoolName" yaml:"defaultBackendAddressPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/application_gateway#default_backend_http_settings_name ApplicationGateway#default_backend_http_settings_name}.
	DefaultBackendHttpSettingsName *string `field:"optional" json:"defaultBackendHttpSettingsName" yaml:"defaultBackendHttpSettingsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/application_gateway#default_redirect_configuration_name ApplicationGateway#default_redirect_configuration_name}.
	DefaultRedirectConfigurationName *string `field:"optional" json:"defaultRedirectConfigurationName" yaml:"defaultRedirectConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/application_gateway#default_rewrite_rule_set_name ApplicationGateway#default_rewrite_rule_set_name}.
	DefaultRewriteRuleSetName *string `field:"optional" json:"defaultRewriteRuleSetName" yaml:"defaultRewriteRuleSetName"`
}

