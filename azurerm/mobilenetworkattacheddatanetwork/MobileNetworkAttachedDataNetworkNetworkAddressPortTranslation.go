// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkattacheddatanetwork


type MobileNetworkAttachedDataNetworkNetworkAddressPortTranslation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/mobile_network_attached_data_network#icmp_pinhole_timeout_in_seconds MobileNetworkAttachedDataNetwork#icmp_pinhole_timeout_in_seconds}.
	IcmpPinholeTimeoutInSeconds *float64 `field:"optional" json:"icmpPinholeTimeoutInSeconds" yaml:"icmpPinholeTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/mobile_network_attached_data_network#pinhole_maximum_number MobileNetworkAttachedDataNetwork#pinhole_maximum_number}.
	PinholeMaximumNumber *float64 `field:"optional" json:"pinholeMaximumNumber" yaml:"pinholeMaximumNumber"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/mobile_network_attached_data_network#port_range MobileNetworkAttachedDataNetwork#port_range}
	PortRange *MobileNetworkAttachedDataNetworkNetworkAddressPortTranslationPortRange `field:"optional" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/mobile_network_attached_data_network#tcp_pinhole_timeout_in_seconds MobileNetworkAttachedDataNetwork#tcp_pinhole_timeout_in_seconds}.
	TcpPinholeTimeoutInSeconds *float64 `field:"optional" json:"tcpPinholeTimeoutInSeconds" yaml:"tcpPinholeTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/mobile_network_attached_data_network#tcp_port_reuse_minimum_hold_time_in_seconds MobileNetworkAttachedDataNetwork#tcp_port_reuse_minimum_hold_time_in_seconds}.
	TcpPortReuseMinimumHoldTimeInSeconds *float64 `field:"optional" json:"tcpPortReuseMinimumHoldTimeInSeconds" yaml:"tcpPortReuseMinimumHoldTimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/mobile_network_attached_data_network#udp_pinhole_timeout_in_seconds MobileNetworkAttachedDataNetwork#udp_pinhole_timeout_in_seconds}.
	UdpPinholeTimeoutInSeconds *float64 `field:"optional" json:"udpPinholeTimeoutInSeconds" yaml:"udpPinholeTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/mobile_network_attached_data_network#udp_port_reuse_minimum_hold_time_in_seconds MobileNetworkAttachedDataNetwork#udp_port_reuse_minimum_hold_time_in_seconds}.
	UdpPortReuseMinimumHoldTimeInSeconds *float64 `field:"optional" json:"udpPortReuseMinimumHoldTimeInSeconds" yaml:"udpPortReuseMinimumHoldTimeInSeconds"`
}

