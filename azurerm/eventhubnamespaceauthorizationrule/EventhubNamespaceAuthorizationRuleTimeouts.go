package eventhubnamespaceauthorizationrule


type EventhubNamespaceAuthorizationRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventhub_namespace_authorization_rule#create EventhubNamespaceAuthorizationRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventhub_namespace_authorization_rule#delete EventhubNamespaceAuthorizationRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventhub_namespace_authorization_rule#read EventhubNamespaceAuthorizationRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventhub_namespace_authorization_rule#update EventhubNamespaceAuthorizationRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
