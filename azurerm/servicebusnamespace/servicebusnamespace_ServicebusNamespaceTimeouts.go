package servicebusnamespace


type ServicebusNamespaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_namespace#create ServicebusNamespace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_namespace#delete ServicebusNamespace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_namespace#read ServicebusNamespace#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_namespace#update ServicebusNamespace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
