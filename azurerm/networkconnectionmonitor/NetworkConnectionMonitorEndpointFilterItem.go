package networkconnectionmonitor


type NetworkConnectionMonitorEndpointFilterItem struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor#address NetworkConnectionMonitor#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor#type NetworkConnectionMonitor#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

