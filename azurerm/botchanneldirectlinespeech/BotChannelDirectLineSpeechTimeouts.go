// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botchanneldirectlinespeech


type BotChannelDirectLineSpeechTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/bot_channel_direct_line_speech#create BotChannelDirectLineSpeech#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/bot_channel_direct_line_speech#delete BotChannelDirectLineSpeech#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/bot_channel_direct_line_speech#read BotChannelDirectLineSpeech#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/bot_channel_direct_line_speech#update BotChannelDirectLineSpeech#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

