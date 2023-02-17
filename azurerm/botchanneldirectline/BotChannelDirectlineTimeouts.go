package botchanneldirectline


type BotChannelDirectlineTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_channel_directline#create BotChannelDirectline#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_channel_directline#delete BotChannelDirectline#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_channel_directline#read BotChannelDirectline#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_channel_directline#update BotChannelDirectline#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

