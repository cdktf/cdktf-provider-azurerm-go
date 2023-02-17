package iothub


type IothubRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#condition Iothub#condition}.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#enabled Iothub#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#endpoint_names Iothub#endpoint_names}.
	EndpointNames *[]*string `field:"optional" json:"endpointNames" yaml:"endpointNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#name Iothub#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#source Iothub#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

