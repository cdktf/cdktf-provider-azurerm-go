package digitaltwinsinstance


type DigitalTwinsInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_instance#create DigitalTwinsInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_instance#delete DigitalTwinsInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_instance#read DigitalTwinsInstance#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_instance#update DigitalTwinsInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

