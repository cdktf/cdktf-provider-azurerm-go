// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionapp


type FunctionAppAuthSettingsGoogle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/function_app#client_id FunctionApp#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/function_app#client_secret FunctionApp#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/function_app#oauth_scopes FunctionApp#oauth_scopes}.
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

