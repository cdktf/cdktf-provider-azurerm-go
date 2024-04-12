// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botchannelwebchat


type BotChannelWebChatSite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/bot_channel_web_chat#name BotChannelWebChat#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/bot_channel_web_chat#endpoint_parameters_enabled BotChannelWebChat#endpoint_parameters_enabled}.
	EndpointParametersEnabled interface{} `field:"optional" json:"endpointParametersEnabled" yaml:"endpointParametersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/bot_channel_web_chat#storage_enabled BotChannelWebChat#storage_enabled}.
	StorageEnabled interface{} `field:"optional" json:"storageEnabled" yaml:"storageEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/bot_channel_web_chat#user_upload_enabled BotChannelWebChat#user_upload_enabled}.
	UserUploadEnabled interface{} `field:"optional" json:"userUploadEnabled" yaml:"userUploadEnabled"`
}

