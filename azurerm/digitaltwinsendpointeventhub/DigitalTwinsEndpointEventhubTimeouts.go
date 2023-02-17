package digitaltwinsendpointeventhub


type DigitalTwinsEndpointEventhubTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_endpoint_eventhub#create DigitalTwinsEndpointEventhub#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_endpoint_eventhub#delete DigitalTwinsEndpointEventhub#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_endpoint_eventhub#read DigitalTwinsEndpointEventhub#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_endpoint_eventhub#update DigitalTwinsEndpointEventhub#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

