package automationconnectiontype


type AutomationConnectionTypeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_connection_type#create AutomationConnectionType#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_connection_type#delete AutomationConnectionType#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_connection_type#read AutomationConnectionType#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
