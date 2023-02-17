package servicebusqueue


type ServicebusQueueTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#create ServicebusQueue#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#delete ServicebusQueue#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#read ServicebusQueue#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#update ServicebusQueue#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

