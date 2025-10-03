// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescalesetpacketcapture


type VirtualMachineScaleSetPacketCaptureFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/virtual_machine_scale_set_packet_capture#protocol VirtualMachineScaleSetPacketCapture#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/virtual_machine_scale_set_packet_capture#local_ip_address VirtualMachineScaleSetPacketCapture#local_ip_address}.
	LocalIpAddress *string `field:"optional" json:"localIpAddress" yaml:"localIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/virtual_machine_scale_set_packet_capture#local_port VirtualMachineScaleSetPacketCapture#local_port}.
	LocalPort *string `field:"optional" json:"localPort" yaml:"localPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/virtual_machine_scale_set_packet_capture#remote_ip_address VirtualMachineScaleSetPacketCapture#remote_ip_address}.
	RemoteIpAddress *string `field:"optional" json:"remoteIpAddress" yaml:"remoteIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/virtual_machine_scale_set_packet_capture#remote_port VirtualMachineScaleSetPacketCapture#remote_port}.
	RemotePort *string `field:"optional" json:"remotePort" yaml:"remotePort"`
}

