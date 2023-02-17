package networkconnectionmonitor


type NetworkConnectionMonitorTestConfigurationSuccessThreshold struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor#checks_failed_percent NetworkConnectionMonitor#checks_failed_percent}.
	ChecksFailedPercent *float64 `field:"optional" json:"checksFailedPercent" yaml:"checksFailedPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor#round_trip_time_ms NetworkConnectionMonitor#round_trip_time_ms}.
	RoundTripTimeMs *float64 `field:"optional" json:"roundTripTimeMs" yaml:"roundTripTimeMs"`
}

