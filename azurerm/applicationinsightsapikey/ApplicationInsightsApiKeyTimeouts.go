package applicationinsightsapikey


type ApplicationInsightsApiKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_insights_api_key#create ApplicationInsightsApiKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_insights_api_key#delete ApplicationInsightsApiKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_insights_api_key#read ApplicationInsightsApiKey#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_insights_api_key#update ApplicationInsightsApiKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

