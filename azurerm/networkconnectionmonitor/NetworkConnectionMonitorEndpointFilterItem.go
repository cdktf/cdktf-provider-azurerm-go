// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectionmonitor


type NetworkConnectionMonitorEndpointFilterItem struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/network_connection_monitor#address NetworkConnectionMonitor#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/network_connection_monitor#type NetworkConnectionMonitor#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

