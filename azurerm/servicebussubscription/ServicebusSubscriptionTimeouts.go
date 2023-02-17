package servicebussubscription


type ServicebusSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#create ServicebusSubscription#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#delete ServicebusSubscription#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#read ServicebusSubscription#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#update ServicebusSubscription#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

