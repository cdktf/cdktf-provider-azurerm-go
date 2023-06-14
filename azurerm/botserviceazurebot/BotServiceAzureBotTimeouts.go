package botserviceazurebot


type BotServiceAzureBotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/bot_service_azure_bot#create BotServiceAzureBot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/bot_service_azure_bot#delete BotServiceAzureBot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/bot_service_azure_bot#read BotServiceAzureBot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/bot_service_azure_bot#update BotServiceAzureBot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

