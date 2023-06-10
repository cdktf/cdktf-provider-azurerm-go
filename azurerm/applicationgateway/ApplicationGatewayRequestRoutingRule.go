package applicationgateway


type ApplicationGatewayRequestRoutingRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#http_listener_name ApplicationGateway#http_listener_name}.
	HttpListenerName *string `field:"required" json:"httpListenerName" yaml:"httpListenerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#rule_type ApplicationGateway#rule_type}.
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#backend_address_pool_name ApplicationGateway#backend_address_pool_name}.
	BackendAddressPoolName *string `field:"optional" json:"backendAddressPoolName" yaml:"backendAddressPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#backend_http_settings_name ApplicationGateway#backend_http_settings_name}.
	BackendHttpSettingsName *string `field:"optional" json:"backendHttpSettingsName" yaml:"backendHttpSettingsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#priority ApplicationGateway#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#redirect_configuration_name ApplicationGateway#redirect_configuration_name}.
	RedirectConfigurationName *string `field:"optional" json:"redirectConfigurationName" yaml:"redirectConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#rewrite_rule_set_name ApplicationGateway#rewrite_rule_set_name}.
	RewriteRuleSetName *string `field:"optional" json:"rewriteRuleSetName" yaml:"rewriteRuleSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/application_gateway#url_path_map_name ApplicationGateway#url_path_map_name}.
	UrlPathMapName *string `field:"optional" json:"urlPathMapName" yaml:"urlPathMapName"`
}

