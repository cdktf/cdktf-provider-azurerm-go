package servicebusqueueauthorizationrule


type ServicebusQueueAuthorizationRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue_authorization_rule#create ServicebusQueueAuthorizationRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue_authorization_rule#delete ServicebusQueueAuthorizationRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue_authorization_rule#read ServicebusQueueAuthorizationRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue_authorization_rule#update ServicebusQueueAuthorizationRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

