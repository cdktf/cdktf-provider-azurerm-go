package iothub


type IothubEnrichment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#endpoint_names Iothub#endpoint_names}.
	EndpointNames *[]*string `field:"optional" json:"endpointNames" yaml:"endpointNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#key Iothub#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#value Iothub#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

