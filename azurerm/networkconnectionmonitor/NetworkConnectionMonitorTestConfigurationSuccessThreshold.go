// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectionmonitor


type NetworkConnectionMonitorTestConfigurationSuccessThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/network_connection_monitor#checks_failed_percent NetworkConnectionMonitor#checks_failed_percent}.
	ChecksFailedPercent *float64 `field:"optional" json:"checksFailedPercent" yaml:"checksFailedPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/network_connection_monitor#round_trip_time_ms NetworkConnectionMonitor#round_trip_time_ms}.
	RoundTripTimeMs *float64 `field:"optional" json:"roundTripTimeMs" yaml:"roundTripTimeMs"`
}

