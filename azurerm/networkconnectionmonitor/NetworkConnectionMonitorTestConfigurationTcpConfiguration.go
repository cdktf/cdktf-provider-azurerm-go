package networkconnectionmonitor


type NetworkConnectionMonitorTestConfigurationTcpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/network_connection_monitor#port NetworkConnectionMonitor#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/network_connection_monitor#destination_port_behavior NetworkConnectionMonitor#destination_port_behavior}.
	DestinationPortBehavior *string `field:"optional" json:"destinationPortBehavior" yaml:"destinationPortBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/network_connection_monitor#trace_route_enabled NetworkConnectionMonitor#trace_route_enabled}.
	TraceRouteEnabled interface{} `field:"optional" json:"traceRouteEnabled" yaml:"traceRouteEnabled"`
}

