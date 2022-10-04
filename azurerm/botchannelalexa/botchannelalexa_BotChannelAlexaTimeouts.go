package botchannelalexa


type BotChannelAlexaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_channel_alexa#create BotChannelAlexa#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_channel_alexa#delete BotChannelAlexa#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_channel_alexa#read BotChannelAlexa#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_channel_alexa#update BotChannelAlexa#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

