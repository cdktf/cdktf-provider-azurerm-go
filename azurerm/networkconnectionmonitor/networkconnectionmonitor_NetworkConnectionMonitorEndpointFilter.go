package networkconnectionmonitor


type NetworkConnectionMonitorEndpointFilter struct {
	// item block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor#item NetworkConnectionMonitor#item}
	Item interface{} `field:"optional" json:"item" yaml:"item"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor#type NetworkConnectionMonitor#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}
