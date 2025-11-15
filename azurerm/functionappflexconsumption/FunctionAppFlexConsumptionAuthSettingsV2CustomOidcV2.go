// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionAuthSettingsV2CustomOidcV2 struct {
	// The ID of the Client to use to authenticate with this Custom OIDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/function_app_flex_consumption#client_id FunctionAppFlexConsumption#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The name of the Custom OIDC Authentication Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/function_app_flex_consumption#name FunctionAppFlexConsumption#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The endpoint that contains all the configuration endpoints for this Custom OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/function_app_flex_consumption#openid_configuration_endpoint FunctionAppFlexConsumption#openid_configuration_endpoint}
	OpenidConfigurationEndpoint *string `field:"required" json:"openidConfigurationEndpoint" yaml:"openidConfigurationEndpoint"`
	// The name of the claim that contains the users name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/function_app_flex_consumption#name_claim_type FunctionAppFlexConsumption#name_claim_type}
	NameClaimType *string `field:"optional" json:"nameClaimType" yaml:"nameClaimType"`
	// The list of the scopes that should be requested while authenticating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/function_app_flex_consumption#scopes FunctionAppFlexConsumption#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

