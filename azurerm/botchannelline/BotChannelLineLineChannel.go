// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botchannelline


type BotChannelLineLineChannel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/bot_channel_line#access_token BotChannelLine#access_token}.
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/bot_channel_line#secret BotChannelLine#secret}.
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

