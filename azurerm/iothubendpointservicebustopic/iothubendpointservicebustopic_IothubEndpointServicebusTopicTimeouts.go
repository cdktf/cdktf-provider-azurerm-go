package iothubendpointservicebustopic


type IothubEndpointServicebusTopicTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub_endpoint_servicebus_topic#create IothubEndpointServicebusTopic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub_endpoint_servicebus_topic#delete IothubEndpointServicebusTopic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub_endpoint_servicebus_topic#read IothubEndpointServicebusTopic#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub_endpoint_servicebus_topic#update IothubEndpointServicebusTopic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
