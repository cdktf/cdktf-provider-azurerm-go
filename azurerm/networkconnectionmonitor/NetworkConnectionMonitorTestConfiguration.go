// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectionmonitor


type NetworkConnectionMonitorTestConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/network_connection_monitor#name NetworkConnectionMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/network_connection_monitor#protocol NetworkConnectionMonitor#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// http_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/network_connection_monitor#http_configuration NetworkConnectionMonitor#http_configuration}
	HttpConfiguration *NetworkConnectionMonitorTestConfigurationHttpConfiguration `field:"optional" json:"httpConfiguration" yaml:"httpConfiguration"`
	// icmp_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/network_connection_monitor#icmp_configuration NetworkConnectionMonitor#icmp_configuration}
	IcmpConfiguration *NetworkConnectionMonitorTestConfigurationIcmpConfiguration `field:"optional" json:"icmpConfiguration" yaml:"icmpConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/network_connection_monitor#preferred_ip_version NetworkConnectionMonitor#preferred_ip_version}.
	PreferredIpVersion *string `field:"optional" json:"preferredIpVersion" yaml:"preferredIpVersion"`
	// success_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/network_connection_monitor#success_threshold NetworkConnectionMonitor#success_threshold}
	SuccessThreshold *NetworkConnectionMonitorTestConfigurationSuccessThreshold `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// tcp_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/network_connection_monitor#tcp_configuration NetworkConnectionMonitor#tcp_configuration}
	TcpConfiguration *NetworkConnectionMonitorTestConfigurationTcpConfiguration `field:"optional" json:"tcpConfiguration" yaml:"tcpConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/network_connection_monitor#test_frequency_in_seconds NetworkConnectionMonitor#test_frequency_in_seconds}.
	TestFrequencyInSeconds *float64 `field:"optional" json:"testFrequencyInSeconds" yaml:"testFrequencyInSeconds"`
}

