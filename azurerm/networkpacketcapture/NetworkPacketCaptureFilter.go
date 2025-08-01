// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpacketcapture


type NetworkPacketCaptureFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_packet_capture#protocol NetworkPacketCapture#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_packet_capture#local_ip_address NetworkPacketCapture#local_ip_address}.
	LocalIpAddress *string `field:"optional" json:"localIpAddress" yaml:"localIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_packet_capture#local_port NetworkPacketCapture#local_port}.
	LocalPort *string `field:"optional" json:"localPort" yaml:"localPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_packet_capture#remote_ip_address NetworkPacketCapture#remote_ip_address}.
	RemoteIpAddress *string `field:"optional" json:"remoteIpAddress" yaml:"remoteIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_packet_capture#remote_port NetworkPacketCapture#remote_port}.
	RemotePort *string `field:"optional" json:"remotePort" yaml:"remotePort"`
}

