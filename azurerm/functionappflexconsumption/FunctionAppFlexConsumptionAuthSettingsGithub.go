// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionAuthSettingsGithub struct {
	// The ID of the GitHub app used for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/function_app_flex_consumption#client_id FunctionAppFlexConsumption#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Client Secret of the GitHub app used for GitHub Login. Cannot be specified with `client_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/function_app_flex_consumption#client_secret FunctionAppFlexConsumption#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The app setting name that contains the `client_secret` value used for GitHub Login. Cannot be specified with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/function_app_flex_consumption#client_secret_setting_name FunctionAppFlexConsumption#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// Specifies a list of OAuth 2.0 scopes that will be requested as part of GitHub Login authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/function_app_flex_consumption#oauth_scopes FunctionAppFlexConsumption#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

