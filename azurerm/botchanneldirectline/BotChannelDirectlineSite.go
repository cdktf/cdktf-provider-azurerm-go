// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botchanneldirectline


type BotChannelDirectlineSite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#name BotChannelDirectline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#enabled BotChannelDirectline#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#endpoint_parameters_enabled BotChannelDirectline#endpoint_parameters_enabled}.
	EndpointParametersEnabled interface{} `field:"optional" json:"endpointParametersEnabled" yaml:"endpointParametersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#enhanced_authentication_enabled BotChannelDirectline#enhanced_authentication_enabled}.
	EnhancedAuthenticationEnabled interface{} `field:"optional" json:"enhancedAuthenticationEnabled" yaml:"enhancedAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#storage_enabled BotChannelDirectline#storage_enabled}.
	StorageEnabled interface{} `field:"optional" json:"storageEnabled" yaml:"storageEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#trusted_origins BotChannelDirectline#trusted_origins}.
	TrustedOrigins *[]*string `field:"optional" json:"trustedOrigins" yaml:"trustedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#user_upload_enabled BotChannelDirectline#user_upload_enabled}.
	UserUploadEnabled interface{} `field:"optional" json:"userUploadEnabled" yaml:"userUploadEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#v1_allowed BotChannelDirectline#v1_allowed}.
	V1Allowed interface{} `field:"optional" json:"v1Allowed" yaml:"v1Allowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/bot_channel_directline#v3_allowed BotChannelDirectline#v3_allowed}.
	V3Allowed interface{} `field:"optional" json:"v3Allowed" yaml:"v3Allowed"`
}

