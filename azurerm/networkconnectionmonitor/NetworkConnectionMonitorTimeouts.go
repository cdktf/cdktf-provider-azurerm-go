package networkconnectionmonitor


type NetworkConnectionMonitorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_connection_monitor#create NetworkConnectionMonitor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_connection_monitor#delete NetworkConnectionMonitor#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_connection_monitor#read NetworkConnectionMonitor#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_connection_monitor#update NetworkConnectionMonitor#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

