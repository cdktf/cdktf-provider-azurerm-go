package networkconnectionmonitor


type NetworkConnectionMonitorTestConfigurationHttpConfigurationRequestHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor#name NetworkConnectionMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor#value NetworkConnectionMonitor#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

