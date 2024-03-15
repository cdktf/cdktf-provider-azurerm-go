// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinepacketcapture


type VirtualMachinePacketCaptureFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/virtual_machine_packet_capture#protocol VirtualMachinePacketCapture#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/virtual_machine_packet_capture#local_ip_address VirtualMachinePacketCapture#local_ip_address}.
	LocalIpAddress *string `field:"optional" json:"localIpAddress" yaml:"localIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/virtual_machine_packet_capture#local_port VirtualMachinePacketCapture#local_port}.
	LocalPort *string `field:"optional" json:"localPort" yaml:"localPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/virtual_machine_packet_capture#remote_ip_address VirtualMachinePacketCapture#remote_ip_address}.
	RemoteIpAddress *string `field:"optional" json:"remoteIpAddress" yaml:"remoteIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/virtual_machine_packet_capture#remote_port VirtualMachinePacketCapture#remote_port}.
	RemotePort *string `field:"optional" json:"remotePort" yaml:"remotePort"`
}

